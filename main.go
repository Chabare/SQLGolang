package main

import (
	"fmt"
	"log"

	"github.com/chabare/sql/sql"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

func createConnection(host, port, user, pass, name string) mysql.Conn {
	d := mysql.New("tcp", "", host+":"+port, user, pass, name)
	err := d.Connect()

	if err != nil {
		log.Fatal(err)
	}

	return d
}

func main() {
	cnff := sql.ConfigFile{"config.ini"}
	cnf := cnff.GetConfig()

	conn := createConnection(cnf.Host, cnf.Port, cnf.User, cnf.Pass, cnf.Name)
	defer conn.Close()
	folder := sql.DBTable{"Folder", []string{"ID", "Name", "Count", "Size", "LastIterationID", "Created"}}

	// values := []string{"asd", "50", "50", "1", "2015-11-29"}
	cnd := []sql.Condition{sql.Condition{"ID", "5"}}

	cols := []string{"Count", "Size"}
	vals := []string{"20", "20"}
	r := folder.Update(&conn, &cols, &vals, &cnd)

	fmt.Println(r.AffectedRows())
}
