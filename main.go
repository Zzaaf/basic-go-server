package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// переменная port используется для указания значения в виде строки
	port := ":3000"
	println("Server listen on port", port)

	// обработчики маршрутов
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/json", jsonResponse)
	http.HandleFunc("/html", htmlResponse)

	// метод http.ListenAndServe запускает HTTP сервер и возвращает ошибку, которую необходимо обработать
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

// функция срабатывающая на запрос по маршруту '/', в качестве ответа строка
func mainPage(w http.ResponseWriter, r *http.Request) {
	welcomeText := "Hello, Go!"
	w.Write([]byte(welcomeText))
}

// описанный тип Message, формиирующий структуру (объект в JS)
type Message struct {
	Language string `json:"language"`
}

// функция срабатывает на запрос по маршруту '/json', в качестве ответа JSON
func jsonResponse(w http.ResponseWriter, r *http.Request) {
	// переменная message формируется на основе типа Message, с переденным значением
	message := Message{"JS"}

	// метод json.Marshal возвращает JSON на основе переданного аргумента
	json, err := json.Marshal(message)

	if err != nil {
		log.Fatal("json.Marshal", err)
	}

	w.Write([]byte(json))
}

func htmlResponse(w http.ResponseWriter, r *http.Request) {
	// метод template.ParseFiles формирует заготовку на основе статического файла
	tmpl, errTemplate := template.ParseFiles("static/index.html")

	if errTemplate != nil {
		http.Error(w, errTemplate.Error(), 400)
		return
	}

	// метод tmpl.Execute формирует ответ сервера в виде HTML-кода
	errHTML := tmpl.Execute(w, nil)
	if errHTML != nil {
		http.Error(w, errTemplate.Error(), 400)
	}
}
