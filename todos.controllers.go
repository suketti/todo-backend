package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TodoEditor struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoItem struct {
	TodoId int    `json:"todoId"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}

func Controller_Todos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := GetTodoItems()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "todoList")

	case http.MethodPost:
		var request TodoEditor
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := CreateTodoItem(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		SendResponse(w, struct{}{})
	}
}

func Controller_Todos_Id(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		result, err := GetTodoItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "todoEditor")

	case http.MethodPatch:
		var request TodoEditor
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateTodoItem(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})

	case http.MethodDelete:
		err := DeleteTodoItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}
