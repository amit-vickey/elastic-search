package main

import (
	"bufio"
	"fmt"
	"os"
	elasticsearch "pkg/elasticsearch.go/pkg"
	"strconv"
	"strings"
)

func main() {
	var localEsUrl = "http://127.0.0.1:9200"
	reader := bufio.NewReader(os.Stdin)
	es, err := elasticsearch.NewESClient(localEsUrl)
	if err != nil {
		fmt.Println("Error while getting ES client ::", err)
		return
	}
	takeAndProcessUserInput(reader, es)
}

func getTrimmedString(str string) string {
	return strings.TrimSuffix(str, "\n")
}

func takeAndProcessUserInput(reader *bufio.Reader, es elasticsearch.IElasticSearchClient) {
	for {
		fmt.Println("1. Check index status")
		fmt.Println("2. Create index")
		fmt.Println("3. Delete Index")
		fmt.Println("4. Quit")
		fmt.Print("Enter your option :: ")
		input, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("Error while reading user input ::", readErr)
			return
		}
		option, intErr := strconv.Atoi(getTrimmedString(input))
		if intErr != nil {
			fmt.Println("Error while reading user input ::", intErr)
			return
		}
		switch option {
		case 1:
			checkIndex(reader, es)
		case 2:
			createIndex(reader, es)
		case 3:
			deleteIndex(reader, es)
		case 4:
			fmt.Println("Bye...")
			return
		}
	}
}

func checkIndex(reader *bufio.Reader, es elasticsearch.IElasticSearchClient) {
	fmt.Print("Enter Index Name :: ")
	index, _ := reader.ReadString('\n')
	index = getTrimmedString(index)
	indexExists, ieErr := es.IndexExists(index)
	if ieErr != nil {
		fmt.Println("Some error occurred ::", ieErr)
		return
	}
	if indexExists {
		fmt.Println("Index exists...")
	} else {
		fmt.Println("Index doesnot exist...")
	}
	fmt.Println()
}

func createIndex(reader *bufio.Reader, es elasticsearch.IElasticSearchClient) {
	fmt.Print("Enter Index Name :: ")
	index, _ := reader.ReadString('\n')
	index = getTrimmedString(index)
	indexExists, ieErr := es.IndexExists(index)
	if ieErr != nil {
		fmt.Println("Some error occurred ::", ieErr)
		return
	}
	if indexExists {
		fmt.Println("Index already exists...")
		return
	}
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
	fmt.Println()
}

func deleteIndex(reader *bufio.Reader, es elasticsearch.IElasticSearchClient) {
	fmt.Print("Enter Index Name :: ")
	index, _ := reader.ReadString('\n')
	index = getTrimmedString(index)
	indexExists, ieErr := es.IndexExists(index)
	if ieErr != nil {
		fmt.Println("Some error occurred ::", ieErr)
		return
	}
	if !indexExists {
		fmt.Println("Index does not exist...")
		return
	}
	deleteRes, deleteErr := es.DeleteIndex(index)
	if deleteErr != nil {
		fmt.Println("Error while deleting the index ::", deleteErr)
		return
	}
	if deleteRes.Acknowledged {
		fmt.Println("Successfully deleted the index...")
	} else {
		fmt.Println("Index not deleted...")
	}
	fmt.Println()
}