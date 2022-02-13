package services

import Models "NoteKeeperAPI/src/database/models"

type AuthService struct{}

func (v AuthService) Register(userData Models.Account) Models.Account {
	return userData
}
