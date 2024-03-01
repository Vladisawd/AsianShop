package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Users struct {
	Id   int
	Name string
}

func handler() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Printf("Сервер работает. Порт:8080")
	err := http.ListenAndServe("localhost:8080", nil) //conf.ServerHost+":"+conf.ServerPort
	if err != nil {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		postUser(w, r)
	default:
		http.Error(w, "Не правильный http метод", http.StatusMethodNotAllowed)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	con := connect(newConf())
	user, err := con.Query(`SELECT "id", "name" FROM "user"`)
	if err != nil {
		fmt.Println(err.Error())
	}

	for user.Next() {
		var u Users
		err = user.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		err := json.NewEncoder(w).Encode(fmt.Sprintf("id: %d, name: %s", u.Id, u.Name))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var newUser Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	con := connect(newConf())
	_, er := con.Exec(fmt.Sprintf(`INSERT INTO "user" ("id","name") VALUES(%d, '%s')`, newUser.Id, newUser.Name))
	if er != nil {
		fmt.Println(er.Error())
	}
	fmt.Fprintf(w, "Успешно добавил нового пользователя: %d, %s)", newUser.Id, newUser.Name)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Сервер работает корректно")
}
