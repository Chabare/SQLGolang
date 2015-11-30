package cSql

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

// SelectTemplate is a template for the Select statement
func SelectTemplate(useConditional bool) *template.Template {
	var tmpl *template.Template
	var err error
	if useConditional {
		tmpl, err = template.ParseFiles("sql/templates/select.tpl", "sql/templates/where.tpl")
	} else {
		tmpl, err = template.ParseFiles("sql/templates/select.tpl")
	}

	if err != nil {
		log.Fatalf("Error in the 'Select' template: %v\n", err)
	}

	return tmpl
}

// InsertTemplate is a template for the Insert statement
func InsertTemplate() *template.Template {
	tmpl, err := template.ParseFiles("sql/templates/insert.tpl")

	if err != nil {
		log.Fatalf("Error in the 'Insert' template: %v\n", err)
	}

	return tmpl
}

// UpdateTemplate is a template for the Insert statement
func UpdateTemplate(useConditional bool) *template.Template {
	var tmpl *template.Template
	var err error
	if useConditional {
		tmpl, err = template.ParseFiles("sql/templates/update.tpl", "sql/templates/where.tpl")
	} else {
		tmpl, err = template.ParseFiles("sql/templates/update.tpl")
	}

	if err != nil {
		log.Fatalf("Error in the 'Update' template: %v\n", err)
	}

	return tmpl
}
