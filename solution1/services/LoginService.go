package services

import "os"

type LoginService interface {
	Login(username string, password string) bool
}
type LoginInformation struct {
	Username string
	Password string
}

func (l LoginInformation) Login(username string, password string) bool {
	//TODO implement me
	if l.Username == username && l.Password == password {
		return true
	}
	return false
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
