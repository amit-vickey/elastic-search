package main

import (
	"bufio"
	"fmt"
	"os"
	elasticsearch "pkg/elasticsearch.go/pkg"
	"strconv"
	"strings"
)

var es elasticsearch.IElasticSearchClient

func main() {
	var localEsUrl = "http://127.0.0.1:9200"
	reader := bufio.NewReader(os.Stdin)
	var err error
	es, err = elasticsearch.NewESClient(localEsUrl)
	if err != nil {
		fmt.Println("Error while getting ES client ::", err)
		return
	}
	takeAndProcessUserInput(reader)
}

func getTrimmedString(str string) string {
	return strings.TrimSuffix(str, "\n")
}

func takeAndProcessUserInput(reader *bufio.Reader) {
	for {
		fmt.Println("1. Check index status")
		fmt.Println("2. Create index")
		fmt.Println("3. Delete Index")
		fmt.Println("4. Work with Student Index")
		fmt.Println("5. Quit")
		fmt.Print("Enter your option :: ")
		input, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("Error while reading user input ::", readErr)
			continue
		}
		option, intErr := strconv.Atoi(getTrimmedString(input))
		if intErr != nil {
			fmt.Println("Error while reading user input ::", intErr)
			continue
		}
		switch option {
		case 1:
			checkIndex(reader)
		case 2:
			createIndex(reader)
		case 3:
			deleteIndex(reader)
		case 4:
			workWithStudentIndex()
		case 5:
			fmt.Println("Bye...")
			return
		default:
			fmt.Println("Invalid input. Try again.")
		}
	}
}

func checkIndex(reader *bufio.Reader) {
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

func createIndex(reader *bufio.Reader) {
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
	createRes, createErr := es.CreateIndex(index, "")
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

func deleteIndex(reader *bufio.Reader) {
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

func workWithStudentIndex() {
	index := "student"
	fmt.Println("INITIATING 'student' INDEX CREATION...")
	isIndexCreationSuccess := createStudentsIndex(index)
	if isIndexCreationSuccess {
		fmt.Println("INSERTING DATA IN INDEX...")
	}
}

func createStudentsIndex(index string) bool {
	isIndexCreationSuccess := false
	indexExists, ieErr := es.IndexExists(index)
	if ieErr != nil {
		fmt.Println("Some error occurred ::", ieErr)
		return isIndexCreationSuccess
	}
	if indexExists {
		fmt.Println("Index already exist. Dropping the index.")
		deleteRes, deleteErr := es.DeleteIndex(index)
		if deleteErr != nil {
			fmt.Println("Error while deleting the index ::", deleteErr)
			return isIndexCreationSuccess
		}
		if deleteRes.Acknowledged {
			fmt.Println("Successfully deleted the index.")
		} else {
			fmt.Println("Index not deleted. Try again...")
			return isIndexCreationSuccess
		}
	}
	createRes, createErr := es.CreateIndex(index, elasticsearch.StudentMapping)
	if createErr != nil {
		fmt.Println("Error while creating the index ::", createErr)
		return isIndexCreationSuccess
	}
	if createRes.Acknowledged {
		fmt.Println("Successfully created the index.")
		isIndexCreationSuccess = true
	} else {
		fmt.Println("Index not created. Try again...")
	}
	return isIndexCreationSuccess
}