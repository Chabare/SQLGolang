package main

import (
	"fmt"

	"github.com/chabare/SQLGolang/cSql"
	_ "github.com/ziutek/mymySql/native" // Native engine
)

func main() {
	cnff := cSql.ConfigFile{"config.ini"}
	cnf := cnff.GetConfig()

	conn := *cSql.CreateConnection(cnf.Host, cnf.Port, cnf.User, cnf.Pass, cnf.Name)
	defer conn.Close()

	folder := cSql.DBTable{"Folder", []string{"ID", "Name", "Count", "Size", "LastIterationID", "Created"}}
	// values := []string{"asd", "50", "50", "1", "2015-11-29"})
	cnd := []cSql.Condition{{"ID", "5"}}

	//cols := []string{"Count", "Size"}
	//vals := []string{"20", "20"}
	r, _ := folder.Select(&conn, &[]string{}, &cnd)

	fmt.Println(r[0].Str(0))
}
