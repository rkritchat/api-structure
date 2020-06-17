package rest

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"golang-structure-api/pkg/adding"
	"golang-structure-api/pkg/getting"
	"net/http"
)

func Handler(get getting.Service, add adding.Service) *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Post("/login", login(get))
	router.Post("/register", register(add))
	return router
}

func login(s getting.Service) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(getting.User)
		if err := json.NewDecoder(r.Body).Decode(&user); err!=nil{
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return
		}

		info, err := s.Login(*user)
		if err!=nil{
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-type", "Application/json")
		json.NewEncoder(w).Encode(info)
	}
}

func register(s adding.Service) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		request := new(adding.User)
		if err := json.NewDecoder(r.Body).Decode(&request); err!=nil{
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return
		}

		if err := s.Register(*request);err!= nil {
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "Application/json")
		json.NewEncoder(w).Encode(request)
	}
}