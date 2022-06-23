package main

import (
	"flag"
	"log"

	"github.com/joaoluizcadore/xmltocsv-go/service"
)

func main() {

	path := flag.String("path", "", "Directory where the XML files are in.")
	templateFile := flag.String("templateFile", "", "Template file contains the struct of columns and xpath")
	outputFile := flag.String("outputFile", "", "The output file")

	flag.Parse()

	if *templateFile == "" {
		log.Fatal("The templateFile argument is missing.")
	}

	if *path == "" {
		log.Fatal("The path argument is missing.")
	}

	if *outputFile == "" {
		log.Fatal("The outputFile argument is missing.")
	}

	service.Run(*templateFile, *path, *outputFile)
}
