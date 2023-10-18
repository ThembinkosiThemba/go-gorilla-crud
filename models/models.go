package models

type User struct {
	Fullname string
	Username string
	Email string
}

type Post struct {
	Title string
	Body string
	Author User
}