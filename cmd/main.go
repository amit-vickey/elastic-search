package main

import (
	"fmt"
	elasticsearch "pkg/elasticsearch.go/pkg"
)

var (
	localEsUrl = "http://127.0.0.1:9200"
	index = "customer123"
)

func main() {
	es, err := elasticsearch.NewESClient(localEsUrl)
	if err != nil {
		fmt.Println("Error while getting ES client ::", err)
		return
	}
	indexExists, ieErr := es.IndexExists(index)
	if ieErr != nil {
		fmt.Println("Some error occurred ::", ieErr)
		return
	}
	if indexExists {
		fmt.Println("Index already exists...")
	} else {
		fmt.Println("Index doesnot exist. Creating index...")
		createRes, createErr := es.CreateIndex(index)
		if createErr != nil {
			fmt.Println("Error while creating the index ::", createErr)
			return
		}
		if createRes.Acknowledged {
			fmt.Println("Successfully created the index...")
		} else {
			fmt.Println("Index not created...")
		}
	}
}
