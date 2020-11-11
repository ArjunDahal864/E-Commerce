package model

// User is Model defination of User
type User struct {
	UUID             string
	Username         string
	Firstname        string
	Lastname         string
	PhoneNumber      string
	Email            string
	Password         string
	ProfilePictureID string
	Admin            bool
	Seller           bool
}
