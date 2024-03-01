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
	conf := newConf()
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Printf("Сервер %s работает. Порт:%s", conf.ServerHost, conf.ServerPort)
	err := http.ListenAndServe(conf.ServerHost+":"+conf.ServerPort, nil)
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
	_, er := con.Exec(fmt.Sprintf(`INSERT INTO "user" ("id","name") VALUES(%d,'%s')`, newUser.Id, newUser.Name))
	if er != nil {
		fmt.Println(er.Error())
	}
	fmt.Fprintf(w, "Успешно добавил нового пользователя: %d, %s)", newUser.Id, newUser.Name)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Сервер работает корректно")
}
