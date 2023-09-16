package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/akmittal/rest_benchmark/go/ent"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := chi.NewRouter()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	r.Get("/movie", getMovies(client))
	r.Post("/movie", createMovie(client))
	http.ListenAndServe(":3000", r)
}

func getMovies(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		limitNum, err := strconv.ParseInt(limit, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies, err := client.Movie.Query().Limit(int(limitNum)).All(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)

	}
}

func createMovie(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movie, err := client.Movie.
			Create().
			SetName("Titanic").
			SetYear(1997).
			SetGenre("Romance").
			SetLanguage("English").
			SetCountry("USA").
			SetActors("Leonardo DiCaprio, Kate Winslet, Billy Zane, Kathy Bates").
			SetDirector("James Cameron").
			Save(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(movie.Name))
	}
}
