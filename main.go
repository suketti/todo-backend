package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SendResponse(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware", r.URL)
			next.ServeHTTP(w, r)
		})
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:qweasd@tcp(localhost:3306)/TODOLIST")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected to MYSQL Server")
	defer db.Close()

	mux := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	mux.HandleFunc("/todos", Controller_Todos).Methods("GET", "POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", Controller_Todos_Id).Methods("GET", "PATCH", "DELETE")

	//mux.Use(Middleware)

	http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(mux))
}
