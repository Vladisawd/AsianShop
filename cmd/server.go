package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Users struct {
	id   int
	name string
}

var user []Users

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

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Пользователь: '%v'", user)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var newUser Users
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	user = append(user, newUser)
	fmt.Fprintf(w, "Успешно добавил нового пользователя: '%v'", newUser)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Сервер работает корректно")
}
