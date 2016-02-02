package main

import (
	"github.com/vinniyo/authCallback"
	"fmt"
	"log"
	"time"

	"github.com/google/google-api-go-client/compute/v1"
)

//scopes https://www.googleapis.com/auth/compute https://www.googleapis.com/auth/devstorage.full_control

func main() {
	date := time.Now().Format("20060102030405")
	snapshot := &compute.Snapshot{
		Name:        "qb" + date,
		Description: "auto backup for quickbooks server",
		SourceDisk:  "quickbooks-sever",
	}

	client, err := authCallback.BuildOAuthHTTPClient()
	if err != nil {
		fmt.Println("Error building OAuth client: %v", err)
	}

	service, err := compute.New(client)
	if err != nil {
		log.Fatalf("Unable to create Compute service: %v", err)
	}

	projectId := "quickbooks-1120"
	res, _ := service.Disks.CreateSnapshot(projectId, "us-central1-b", "quickbooks-sever", snapshot).Do()
	fmt.Println(res)

}
