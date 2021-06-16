package models

import "github.com/google/uuid"

type Goal struct {
	ID         uuid.UUID `db:"id" json:"id"`
	UserID     uuid.UUID `db:"user_id" json:"userId"`
	IsComplete int       `db:"is_complete" json:"isComplete"`
	Goal       string    `db:"goal" json:"goal"`
}

type GoalRepo interface {
	GetOneGoal(id uuid.UUID) (Goal, error)
	GetGoals() ([]Goal, error)
	CreateGoal(g *Goal) error
	UpdateGoal(g *Goal) error
	DeleteGoal(id uuid.UUID) error
}

type Repo interface {
	GoalRepo
}
