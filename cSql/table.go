package cSql

import (
	"bytes"
	"log"

	"github.com/ziutek/mymysql/mysql"
)

// DBTable Describes a database table
type DBTable struct {
	Name    string
	Columns []string
}

// Select Selects values from table
func (table *DBTable) Select(conn *mysql.Conn, columns *[]string, conditions *[]Condition) ([]mysql.Row, mysql.Result) {
	if len(*columns) == 0 {
		*columns = table.Columns
	}
	// Create the struct for the 'Select'-template
	sStruct := SelectStruct{table.Name, *columns, *conditions}
	// Get the 'Select'-template
	sTemplate := SelectTemplate(len(*conditions) > 0)

	var sql bytes.Buffer
	// Fill the struct values into the template
	sTemplate.Execute(&sql, sStruct)

	// Execute the query
	rows, res, err := mysql.Query(*conn, sql.String())

	if err != nil {
		log.Fatal(err)
	}

	return rows, res
}

// Insert Inserts values into table
func (table *DBTable) Insert(conn *mysql.Conn, columns *[]string, values *[]string) mysql.Result {
	// Check for empty column list
	if len(*columns) == 0 {
		*columns = table.Columns
	}

	// Check for discarded auto-increment value and adjust
	cols := *columns
	if len(*columns)-1 == len(*values) {
		cols = cols[1:]
	} else if len(cols) != len(*values) {
		log.Fatalf("Invalid argument list: Column: %v\nValues: %v\n", cols, values)
	}

	// Create the struct for the 'Insert'-template
	iStruct := InsertStruct{table.Name, cols, *values}
	// Get the 'Insert'-template
	iTemplate := InsertTemplate()

	var sql bytes.Buffer
	// Fill the struct values into the template
	iTemplate.Execute(&sql, iStruct)

	// Execute the query
	_, res, err := mysql.Query(*conn, sql.String())

	if err != nil {
		log.Fatal(err)
	}

	return res
}

// Update Updates a row in the table
func (table *DBTable) Update(conn *mysql.Conn, columns *[]string, values *[]string, conditions *[]Condition) mysql.Result {
	// Check for empty column list
	if len(*columns) == 0 {
		*columns = table.Columns
	}

	// Check for discarded auto-increment value and adjust
	cols := *columns
	if len(*columns)-1 == len(*values) {
		cols = cols[1:]
	} else if len(cols) != len(*values) {
		log.Fatalf("Invalid argument list: Column: %v\nValues: %v\n", cols, values)
	}

	// Create the struct for the 'Update'-template
	uStruct := UpdateStruct{table.Name, cols, *values, *conditions}
	// Get the 'Update'-template
	uTemplate := UpdateTemplate(len(*conditions) > 0)

	var sql bytes.Buffer
	// Fill the struct values into the template
	uTemplate.Execute(&sql, uStruct)

	// Execute the query
	_, res, err := mysql.Query(*conn, sql.String())

	if err != nil {
		log.Fatal(err)
	}

	return res
}
