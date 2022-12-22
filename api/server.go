package api

import (
	"WebApp/database"
	"WebApp/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	database *database.Database
}

func NewServer(database *database.Database) *Server {
	return &Server{database: database}
}

func (s *Server) StartServer() {

	router := mux.NewRouter()

	router.HandleFunc("/", s.Hello).Methods("GET")

	router.HandleFunc("/user", s.NewUser).Methods("POST")

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

func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	//output := s.database.SelectAllTable()

	userService := users.NewUserService(s.database)

	userService.CreateUser()

	json.NewEncoder(w).Encode("new user")
}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	output := s.database.SelectAllTable()
	json.NewEncoder(w).Encode(output)
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
