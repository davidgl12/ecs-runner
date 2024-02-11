package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := ecs.NewFromConfig(cfg)

	list, err := client.ListClusters(context.TODO(), &ecs.ListClustersInput{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(list.ClusterArns)

	fmt.Print("Holaaa")
}
