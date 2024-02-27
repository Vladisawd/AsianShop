package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgBase     string
}

var cfg setting

func init() {
	file, err := os.Open("setting.cfg")
	if err != nil {
		fmt.Println(err.Error())

		panic("Не удалось открыть файл")
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())

		panic("Не удалось прочитать информацию о файле")
	}

	fileByte := make([]byte, stat.Size())

	_, err = file.Read(fileByte)
	if err != nil {
		fmt.Println(err.Error())

		panic("Не удалось прочитать файл конфигурации")
	}

	err = json.Unmarshal(fileByte, &cfg)
	if err != nil {
		fmt.Println(err.Error())

		panic("Не считать данные")
	}
}
