package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	zksStr := "localhost:2181"
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	return conn
}

func main() {
	conn := connect()
	data, _, _ := conn.Children("/brokers/topics/test/partitions")
	fmt.Println("Total Partition Count", len(data))
	for _, partitions := range data {
		fmt.Println(partitions)
	}
}
