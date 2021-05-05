package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var First_Time = true

var t, t1 int64

var lag, totallag int64

type offset struct {
	PartitionOffsets       []int64 `json:"partitionOffsets,omitempty"`
	PartitionLatestOffsets []int64 `json:"partitionLatestOffsets,omitempty"`
}

var objmap map[string]offset

type Topics struct {
	Topics []string `json:"topics,omitempty"`
}

type Consumers struct {
	Consumers []struct {
		Name string `json:"name,omitempty"`
	} `json:"consumers,omitempty"`
}

var consumerlist Consumers
var topiclist Topics

func main() {
	var provider, env, cluster_name string
	var help bool

	flag.StringVar(&provider, "provider", "aws", "it should be either aws or gcp or azure")
	flag.StringVar(&env, "env", "staging", "it should be either production, staging")
	flag.StringVar(&cluster_name, "cluster_name", "kafka-staging", "it should be kakfa cluster name")
	flag.BoolVar(&help, "help", false, "Display Help")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	environment := provider + "-" + env
	details := map[string]map[string]string{
		"aws-production": {"url": "kafka-manager-production.gh.pvt"},
		"aws-staging":    {"url": "kafka-manager-staging.gh.pvt"}
	}

	filename := fmt.Sprintf("%s-%s.json", provider, env)

	file, err := ioutil.ReadFile(filename)

	checkError(err)

	if err := json.Unmarshal(file, &topiclist); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	x := get_consumer_group_mapping(details[environment]["url"], cluster_name, topiclist.Topics)

	for true {
		if First_Time {
			t = time.Now().Unix()
			fmt.Println(t)
			time.Sleep(60 * time.Second)
			fmt.Println("True")

		} else {
			fmt.Println("False")
			t1 = time.Now().Unix()
			diff := t1 - t
			t = t1
			for topic, consumer_group := range x {
				url := fmt.Sprintf("http://%s/api/status/%s/%s/%s/groupSummary", details[environment]["url"], cluster_name, consumer_group, "ZK")
				resp, err := http.Get(url)
				checkError(err)

				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					log.Fatalln("Error Status not OK:", resp.StatusCode)
				}

				body, err := ioutil.ReadAll(resp.Body)
				checkError(err)

				if err := json.Unmarshal(body, &objmap); err != nil {
					log.Fatal(err)
				}

				partitionOffsets := objmap[topic].PartitionOffsets
				partitionLatestOffsets := objmap[topic].PartitionLatestOffsets
				fmt.Println(partitionOffsets, partitionLatestOffsets)

				for i, _ := range partitionOffsets {
					lag = (partitionLatestOffsets[i] - partitionOffsets[i]) / diff
					totallag = totallag + lag
				}
				out := fmt.Sprintf("total lag for %s is %d", consumer_group, totallag)
				fmt.Println(out)
			}
			time.Sleep(60 * time.Second)
		}
		First_Time = false
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func get_consumer_group_mapping(url string, name string, topics []string) map[string]string {
	list_of_consumers := []string{}
	topic_cgs_mapping := make(map[string]string)
	consumer_group_url := fmt.Sprintf("http://%s/api/status/%s/consumersSummary", url, name)
	fmt.Println(consumer_group_url)

	resp, err := http.Get(consumer_group_url)
	checkError(err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	if err := json.Unmarshal(body, &consumerlist); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	for _, v := range consumerlist.Consumers {
		list_of_consumers = append(list_of_consumers, v.Name)
	}

	for _, v := range topics {
		topic_cgs_mapping[v] = fmt.Sprintf("urep-%s", v)
	}

	return topic_cgs_mapping
}
