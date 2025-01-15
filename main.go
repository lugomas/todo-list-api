package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", Signup)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/to-dos", TodosHandler)
	r.HandleFunc("/to-dos/{id}", TodoHandler)
	http.Handle("/", r)
}

func Signup(w http.ResponseWriter, r *http.Request)       {}
func Login(w http.ResponseWriter, r *http.Request)        {}
func TodosHandler(w http.ResponseWriter, r *http.Request) {}
func TodoHandler(w http.ResponseWriter, r *http.Request)  {}
