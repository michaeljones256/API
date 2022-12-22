package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) StartServer() {

	router := mux.NewRouter()

	router.HandleFunc("/", s.Hello).Methods("GET")

	// Get all movies
	router.HandleFunc("/data", s.Get).Methods("GET")

	// Create a movie
	router.HandleFunc("/data", s.Create).Methods("POST")

	// Delete a specific movie by the movieID
	router.HandleFunc("/data/{id}", s.Delete).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/data", s.DeleteAll).Methods("DELETE")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello movies")
}
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Post(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID := params["id"]
	fmt.Println(ID)

}

func (s *Server) DeleteAll(w http.ResponseWriter, r *http.Request) {

}
