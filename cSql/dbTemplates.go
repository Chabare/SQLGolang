package cSql

import (
	"log"
	"path"
	"runtime"
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

var _, base, _, _ = runtime.Caller(1)

// SelectTemplate is a template for the Select statement
func SelectTemplate(useConditional bool) *template.Template {
	var tmpl *template.Template
	var err error

	selectTpl := path.Join(path.Dir(base), "/templates/select.tpl")
	if useConditional {
		whereTpl := path.Join(path.Dir(base), "/templates/where.tpl")
		tmpl, err = template.ParseFiles(selectTpl, whereTpl)
	} else {
		tmpl, err = template.ParseFiles(selectTpl)
	}

	if err != nil {
		log.Fatalf("Error in the 'Select' template: %v\n", err)
	}

	return tmpl
}

// InsertTemplate is a template for the Insert statement
func InsertTemplate() *template.Template {
	insertTpl := path.Join(path.Dir(base), "/templates/insert.tpl")
	tmpl, err := template.ParseFiles(insertTpl)

	if err != nil {
		log.Fatalf("Error in the 'Insert' template: %v\n", err)
	}

	return tmpl
}

// UpdateTemplate is a template for the Insert statement
func UpdateTemplate(useConditional bool) *template.Template {
	var tmpl *template.Template
	var err error
	updateTpl := path.Join(path.Dir(base), "/templates/update.tpl")
	if useConditional {
		whereTpl := path.Join(path.Dir(base), "/templates/where.tpl")
		tmpl, err = template.ParseFiles(updateTpl, whereTpl)
	} else {
		tmpl, err = template.ParseFiles(updateTpl)
	}

	if err != nil {
		log.Fatalf("Error in the 'Update' template: %v\n", err)
	}

	return tmpl
}
