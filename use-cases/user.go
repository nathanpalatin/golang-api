package usecase

import "apigo/model"

type User struct {

}

func NewUserUseCase() User {
	return User{}
}

func (u *User) GetUsers()  ([]model.User, error){
	return []model.User{}, nil
}