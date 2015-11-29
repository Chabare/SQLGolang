package sql

import (
	"log"
	"text/template"
)

// SelectStruct is a struct to use with the SelectTemplate
type SelectStruct struct {
	Table      string
	Columns    []string
	Conditions []Condition
}

// InsertStruct is a struct to use with the InsertTemplate
type InsertStruct struct {
	Table   string
	Columns []string
	Values  []string
}

// UpdateStruct is a struct to use with the UpdateTemplate
type UpdateStruct struct {
	Table      string
	Columns    []string
	Values     []string
	Conditions []Condition
}

var whereTemplate = " WHERE {{with .Conditions}}{{range $condIndex, $cond := .}}{{if $condIndex}} AND {{end}}`{{$cond.Column}}`='{{$cond.Value}}'{{end}}{{end}}"

// SelectTemplate is a template for the Select statement
func SelectTemplate(useConditional bool) *template.Template {
	tmpl := template.New("selectTemplate")

	tmplString := "SELECT {{range $colIndex, $column := .Columns}}{{if $colIndex}}, {{end}}{{if eq $column \"*\"}}{{$column}}{{else}}`{{$column}}`{{end}}{{end}}"
	tmplString += " FROM `{{.Table}}`"
	if useConditional {
		tmplString += whereTemplate
	}
	tmpl, err := tmpl.Parse(tmplString)

	if err != nil {
		log.Fatalf("Error in the 'Select' template: %v\n", err)
	}

	return tmpl
}

// InsertTemplate is a template for the Insert statement
func InsertTemplate() *template.Template {
	tmpl := template.New("insertTemplate")

	tmplString := "INSERT INTO `{{.Table}}`({{range $index, $column := .Columns}}{{if $index}}, {{end}}`{{$column}}`{{end}})"
	tmplString += " VALUES ({{range $index, $value := .Values}}{{if $index}}, {{end}}'{{$value}}'{{end}})"
	tmpl, err := tmpl.Parse(tmplString)

	if err != nil {
		log.Fatalf("Error in the 'Insert' template: %v\n", err)
	}

	return tmpl
}

// UpdateTemplate is a template for the Insert statement
func UpdateTemplate(useConditional bool) *template.Template {
	tmpl := template.New("updateTemplate")

	tmplString := "UPDATE `{{.Table}}` SET "
	tmplString += "{{range $colIndex, $column := .Columns}}{{if $colIndex}}, {{end}}`{{$column}}`='{{index $.Values $colIndex}}'{{end}}"
	if useConditional {
		tmplString += whereTemplate
	}
	tmpl, err := tmpl.Parse(tmplString)

	if err != nil {
		log.Fatalf("Error in the 'Update' template: %v\n", err)
	}

	return tmpl
}
