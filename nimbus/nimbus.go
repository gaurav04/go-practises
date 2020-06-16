package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Topologies struct {
	Topologies []Topology `json:"topologies"`
}

type Topology struct {
	Name string `json:"name"`
}

type Supervisors struct {
	Supervisors []Supervisor `json:"supervisors"`
}

type Supervisor struct {
	Host string `json:"host"`
}

type Cluster struct {
	SlotsFree  int `json:"slotsFree"`
	SlotsTotal int `json:"slotsTotal"`
}

func getTopologies() {

	resp, err := http.Get("http://nimbus:8080/api/v1/topology/summary")

	if err != nil {
		fmt.Println(err)
	}

	databytes, err := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(databytes))

	var topologies Topologies

	json.Unmarshal(databytes, &topologies)

	// Topolgies list
	for i := 0; i < len(topologies.Topologies); i++ {
		fmt.Println(topologies.Topologies[i].Name)
	}

	defer resp.Body.Close()
}

func getSupervisorList() {

	resp, err := http.Get("http://nimbus:8080/api/v1/supervisor/summary")

	if err != nil {
		fmt.Println(err)
	}

	databytes, err := ioutil.ReadAll(resp.Body)
	var supervisors Supervisors

	json.Unmarshal(databytes, &supervisors)

	// Supervisor list
	for i := 0; i < len(supervisors.Supervisors); i++ {
		fmt.Println(supervisors.Supervisors[i].Host)
	}
}

func getCluster() {

	resp, err := http.Get("http://nimbus:8080/api/v1/cluster/summary")

	if err != nil {
		fmt.Println(err)
	}

	databytes, err := ioutil.ReadAll(resp.Body)
	var cluster Cluster

	json.Unmarshal(databytes, &cluster)

	fmt.Println("free slots:", cluster.SlotsFree)

	fmt.Println("total slots:", cluster.SlotsTotal)

}
func main() {

	getTopologies()
	fmt.Println("Supervisors List")
	getSupervisorList()
	fmt.Println("Cluster Info")
	getCluster()

}
