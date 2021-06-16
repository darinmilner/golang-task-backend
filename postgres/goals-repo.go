package postgres

import (
	"fmt"
	"log"
	"todos/api/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GoalsRepo struct {
	*sqlx.DB
}

func (s *GoalsRepo) GetOneGoal(id uuid.UUID) (models.Goal, error) {

	var g models.Goal

	if err := s.Get(&g, "SELECT * FROM goals WHERE id = $1", id); err != nil {
		return models.Goal{}, fmt.Errorf("error getting a goal: %w", err)
	}
	return g, nil
}
func (s *GoalsRepo) GetGoals() ([]models.Goal, error) {
	var ts []models.Goal
	if err := s.Select(&ts, "SELECT * FROM goals"); err != nil {
		return []models.Goal{}, fmt.Errorf("error getting all goals: %w", err)
	}
	log.Print(ts)
	return ts, nil

}

func (s *GoalsRepo) CreateGoal(g *models.Goal) error {
	if err := s.Get(g, "INSERT INTO goals VALUES ($1, $2, $3, $4) RETURNING *", g.ID, g.UserID, g.IsComplete, g.Goal); err != nil {
		return fmt.Errorf("error creating goal: %w", err)
	}
	return nil
}

func (s *GoalsRepo) UpdateGoal(g *models.Goal) error {
	if err := s.Get(g, "UPDATE goals SET is_complete = $1, description = $2, WHERE id = $3 RETURNING *",
		g.IsComplete, g.Goal, g.ID); err != nil {
		return fmt.Errorf("error updating goal: %w", err)
	}
	return nil
}

func (s *GoalsRepo) DeleteGoal(id uuid.UUID) error {
	if _, err := s.Exec("DELETE FROM goals WHERE id = $1", id); err != nil {
		return fmt.Errorf("error deleting goal: %w", err)
	}
	return nil
}
