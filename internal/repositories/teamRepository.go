package repositories

import "racer/form/internal/models"

type TeamRepository interface {
	SaveTeamName(text string) (*models.Team, error)
}

func NewTeamRepository() *models.Team {
	return &models.Team{}
}

func SaveTeamName(text string) (*models.Team, error) {
	return &models.Team{
		TeamName: text,
	}, nil
}
