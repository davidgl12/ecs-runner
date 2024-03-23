package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ecs-runner/internal/model"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	tea "github.com/charmbracelet/bubbletea"
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

	p := tea.NewProgram(model.GetInitialModel(list.ClusterArns))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	fmt.Println(list.ClusterArns)

	fmt.Print("Holaaa")
}
