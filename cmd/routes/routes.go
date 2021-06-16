package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todos/api/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

type Handler struct {
	*chi.Mux
	repo models.GoalRepo
}

func NewServer(repo models.GoalRepo) *Handler {
	mux := &Handler{
		Mux:  chi.NewMux(),
		repo: repo,
	}

	mux.Use(middleware.Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	mux.Use(cors.Handler)
	mux.Route("/goals", func(r chi.Router) {
		r.Get("/", mux.getGoals)
		r.Get("/{id}", mux.getOneGoal)
		r.Post("/create", mux.createGoal)
		r.Post("/{id}/delete", mux.deleteGoal)
		r.Patch("/update/{id}", mux.updateGoal)
	})
	return mux
}

func (s *Handler) getGoals(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Goals []models.Goal
	}
	goals, err := s.repo.GetGoals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(goals)
}

func (s *Handler) getOneGoal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	goal, err := s.repo.GetOneGoal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(goal)
}

type Goal struct{}

func (s *Handler) createGoal(w http.ResponseWriter, r *http.Request) {
	userId := uuid.New()
	isComplete := 0
	headerContentType := r.Header.Get("Content-Type")
	log.Println(headerContentType)

	//g := r.Context().Value(models.Goal{}).(models.Goal)

	//log.Print(g)
	reqBody, _ := ioutil.ReadAll(r.Body)

	type postData struct {
		goal string
	}

	//goal := string(reqBody)
	var data models.Goal
	json.Unmarshal(reqBody, &data)
	//log.Print(string(reqBody))
	err := s.repo.CreateGoal(&models.Goal{
		ID:         uuid.New(),
		UserID:     userId,
		IsComplete: isComplete,
		//Goal:       data.goal,
		Goal: data.Goal,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorJSON{
			Msg:        "Goal Does not exist",
			ErrMsg:     err,
			StatusCode: http.StatusNotFound,
		})
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Handler) deleteGoal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.repo.DeleteGoal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Handler) updateGoal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	type postData struct {
		goal string
	}

	//goal := string(reqBody)
	var data models.Goal
	json.Unmarshal(reqBody, &data)
	err = s.repo.UpdateGoal(&models.Goal{
		ID:         id,
		IsComplete: data.IsComplete,
		Goal:       data.Goal,
	})
}
