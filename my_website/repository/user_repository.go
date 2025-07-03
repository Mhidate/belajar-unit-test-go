package repository

import "github.com/Mhidate/belajar-unit-test-go/my_website/model"

type UserRepository interface {
	GetByID(id int) (*model.User, error)
}
