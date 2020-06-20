package main

import (
	"fmt"
	"log"

	marathon "github.com/gambol99/go-marathon"
)

func getMarathonServiceNames(client marathon.Marathon) ([]string, error) {
	return client.ListApplications(nil)
}

func listMarathonServices(client marathon.Marathon) {
	applications, err := client.Applications(nil)
	if err != nil {
		log.Fatalf("Failed to list applications")
	}
	fmt.Printf("Found %d applications running\n", len(applications.Apps))
	for _, application := range applications.Apps {
		details, err := client.Application(application.ID)
		if err != nil {
			panic(err)
		}
		if details.Tasks != nil && len(details.Tasks) > 0 {
			health, err := client.ApplicationOK(details.ID)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Application: %s, healthy: %t, instances: %d, cpus: %f\n", details.ID, health, len(details.Tasks), details.CPUs*float64(len(details.Tasks)))
			for _, task := range details.Tasks {
				if len(task.Ports) > 0 {
					fmt.Printf("  http://%s:%d\n", task.Host, task.Ports[0])
				} else {
					fmt.Println("No Port Assigned")
				}
			}
		}
	}
}

func main() {

	marathonURL := "http://<hostname>:8080"
	config := marathon.NewDefaultConfig()
	config.URL = marathonURL
	config.HTTPBasicAuthUser = "username"
	config.HTTPBasicPassword = "password"
	client, err := marathon.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}

	services, _ := getMarathonServiceNames(client)
	fmt.Println(services)
	for _, v := range services {
		fmt.Println(v)
	}

	listMarathonServices(client)
}
