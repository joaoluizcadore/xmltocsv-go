package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joaoluizcadore/xmltocsv-go/util"
)

type templateRow struct {
	columnName    string
	xpath         string
	valueLocation string
}
type Template struct {
	rows []templateRow
}

func LoadTemplate(templateFile string) (*Template, error) {
	file, err := os.Open(templateFile)
	util.CheckError(err, "Cannot open the file template.")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	template := &Template{
		rows: make([]templateRow, 0),
	}

	for scanner.Scan() {
		template.rows = append(template.rows, getTemplateRow(scanner.Text()))
	}

	if len(template.rows) == 0 {
		return nil, errors.New("The template file is empty.")
	}

	return template, nil
}

func validateTemplateRow(text string) error {
	columns := strings.Split(text, ";")
	if len(columns) != 3 {
		return errors.New("The template line needs to have 3 parameters (column name;xpath;value location)")
	}

	valueLocation := columns[2]
	if valueLocation != "ATTR" && valueLocation != "TEXT" {
		return errors.New("The value location should be ATTR for attribute values or TEXT for field values.")
	}

	return nil
}

func getTemplateRow(text string) templateRow {
	err := validateTemplateRow(text)
	util.CheckError(err, fmt.Sprintf("Invalid template line: %v", text))

	columns := strings.Split(text, ";")
	columnName := columns[0]
	xpath := columns[1]
	valueLocation := columns[2]

	row := &templateRow{
		columnName:    columnName,
		xpath:         xpath,
		valueLocation: valueLocation,
	}

	return *row
}
