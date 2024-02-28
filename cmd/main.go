package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	conf := newConf()

	con := connect(conf)

	_, err := con.Exec(`INSERT INTO "user" ("name") VALUES('Aboba')`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
