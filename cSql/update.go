package cSql

import (
	"bytes"

	"github.com/ziutek/mymysql/mysql"
)

// Update updates a row
func Update(conn *mysql.Conn, table string, columns, values *[]string, conditions *[]Condition) mysql.Result {
	uStruct := UpdateStruct{table, *columns, *values, *conditions}
	uTemplate := UpdateTemplate(len(*conditions) > 0)

	var sql bytes.Buffer
	uTemplate.Execute(&sql, uStruct)

	_, res, _ := mysql.Query(*conn, sql.String())

	return res
}
