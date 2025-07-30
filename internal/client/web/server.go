package web

import (
	_ "embed"
	"html/template"
	"net/http"
	"reflect"
)

//go:embed index.html
var indexHTML []byte

var board *Board

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		board = NewBoard(6)
		Render(w)
	})

	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		board.Start()
		Render(w)
	})

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		board.Stop()
		Render(w)
	})

	mux.HandleFunc("/right", func(w http.ResponseWriter, r *http.Request) {
		board.Right()
		Render(w)
	})

	mux.HandleFunc("/left", func(w http.ResponseWriter, r *http.Request) {
		board.Left()
		Render(w)
	})

	mux.HandleFunc("/forward", func(w http.ResponseWriter, r *http.Request) {
		board.Forward()
		Render(w)
	})

	mux.HandleFunc("/back", func(w http.ResponseWriter, r *http.Request) {
		board.Back()
		Render(w)
	})
	return mux
}

func Server() {
	mux := SetupRoutes()
	http.ListenAndServe(":8080", mux)
}

func eq(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func Render(w http.ResponseWriter) {
	board.Print()
	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{"eq": eq}).Parse(string(indexHTML)))
	tmpl.Execute(w, board)
}