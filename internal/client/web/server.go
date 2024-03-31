package web

import (
	"html/template"
	"net/http"
	"reflect"
)

var board *Board

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		board = NewBoard(6)
		Render(w)
	})

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		board.Start()
		Render(w)
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		board.Stop()
		Render(w)
	})

	http.HandleFunc("/right", func(w http.ResponseWriter, r *http.Request) {
		board.Right()
		Render(w)
	})

	http.HandleFunc("/left", func(w http.ResponseWriter, r *http.Request) {
		board.Left()
		Render(w)
	})

	http.HandleFunc("/forward", func(w http.ResponseWriter, r *http.Request) {
		board.Forward()
		Render(w)
	})

	http.HandleFunc("/back", func(w http.ResponseWriter, r *http.Request) {
		board.Back()
		Render(w)
	})

	http.ListenAndServe(":8080", nil)
}

func eq(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func Render(w http.ResponseWriter) {
	board.Print()
	tmpl := template.Must(template.ParseFiles("./internal/client/web/index.html"))
	tmpl.Funcs(template.FuncMap{"eq": eq})
	tmpl.Execute(w, board)
}
