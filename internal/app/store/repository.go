package store

import "github.com/israpilovsha/Rest-Api-Task/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
