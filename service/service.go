package service

import (
	"fmt"

	"github.com/joaoluizcadore/xmltocsv-go/util"
)

func Run(templateFile string, path string, outputFile string) {
	template, err := LoadTemplate(templateFile)
	util.CheckError(err, "Error to load the template file.")

	fmt.Println(template)
}
