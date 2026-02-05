package store

import "learn-go/No8/model"

func CreatUser(username, password string) *model.User {
	u := &model.User{
		Username: username,
		Password: password,
	}
	return u
}

func GetUserByName(username string) *model.User {
	//TODO Query DB
	return &model.User{}
}
