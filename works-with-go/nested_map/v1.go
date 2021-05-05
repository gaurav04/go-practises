package main

import "fmt"

func main() {
	shoppingList := map[string]map[string]int{"aws-production": {"url": "kafka-manager.production-default-uw1.mistk8s.pvt"},
	"aws-staging": {"url":"kafka-manager.staging-default-ue1.mistk8s.pvt"},
	"gcp-production": {"url": "kafka-manager.mist.pvt"},
	"gcp-staging": {"url":"kafka-manager.mist.pvt"},
	"aws-eu": {"url":"kafka-manager-eu.mist.pvt"}
 }

}
