package main

import (
	"fmt"
	. "code.cloudfoundry.org/cli/cf/manifest"
	"os"
	"log"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please specify the path to the child manifest you want to validate.")
		os.Exit(1)
	}

	var repo Repository

	repo = NewDiskRepository()
	path := args[0]
	m, err := repo.ReadManifest(path)

	if err != nil {
		log.Fatalf("Error while reading manifest file via cf cli: \n %s", err)
		os.Exit(1)
	}

	log.Println("cf cli successfully read deploy manifest")
	log.Println(m.Data)
	os.Exit(0)

}