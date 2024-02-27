package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	con := connect()
	if con != nil {
		fmt.Println(con.Error())
		return
	}
	fmt.Println("Ready!")

	_, err := db.Exec(`INSERT INTO "user" ("name") VALUES('Aboba')`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
