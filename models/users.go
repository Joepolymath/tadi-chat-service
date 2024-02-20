package models

import "time"

type User struct {
	ID             string    `bson:"_id" json:"_id"`
	FirstName      string    `bson:"firstName" json:"firstName"`
	LastName       string    `bson:"lastName" json:"lastName"`
	Title          string    `bson:"title" json:"title"`
	Email          string    `bson:"email" json:"email"`
	Phone          string    `bson:"phone" json:"phone"`
	Role           string    `bson:"role" json:"role"`
	ProfilePicture string    `bson:"profilePicture" json:"profilePicture"`
	Password       string    `bson:"password" json:"password"`
	DateOfBirth    string    `bson:"dateOfBirth" json:"dateOfBirth"`
	Address        string    `bson:"address" json:"address"`
	Nationality    string    `bson:"nationality" json:"nationality"`
	Username       string    `bson:"username" json:"username"`
	Gender         string    `bson:"gender" json:"gender"`
	KnownIps       []string  `bson:"knownIps" json:"knownIps"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
	FlaggedIP      string    `bson:"flaggedIp" json:"flaggedIp"`
	LastSeen      time.Time    `bson:"lastSeen" json:"lastSeen"`
}
