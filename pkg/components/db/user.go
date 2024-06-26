package db

import (
	"paragrafio/pkg/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	*TransactionProvider
}

func (r *UserRepository) IsEmailConfirmed(user models.User) bool {
	return false
}

func (r *UserRepository) FindByID(transaction models.Transaction, id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(transaction models.Transaction, email string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByCredentials(transaction models.Transaction, email string, password string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) Insert(transaction models.Transaction, newUser models.NewUser) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) Patch(transaction models.Transaction, patchedUser models.PatchedUser, previousUser models.User) (*models.User, error) {
	return nil, nil
}