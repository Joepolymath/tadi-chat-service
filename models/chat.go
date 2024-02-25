package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat instance
type Chat struct {
	ChatName    string    `bson:"chatName" json:"chatName"`
	IsGroupChat bool      `bson:"isGroupChat" json:"isGroupChat"`
	UsersIds       []primitive.ObjectID  `bson:"usersIds" json:"usersIds"`
	Users       []*User  `bson:"users" json:"users"`
	GroupAdmins  []primitive.ObjectID    `bson:"groupAdmins" json:"groupAdmin"`
	LastMessage *Message		`bson:"lastMessage" json:"lastMessage"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt" json:"updatedAt"`
}

// Message in each chat
type Message struct {
	Sender  string    `json:"sender" bson:"sender"`
	Content string    `json:"content" bson:"content"`
	ChatID    primitive.ObjectID      `json:"chatId" bson:"chatId"`
	Chat    *Chat      `json:"chat" bson:"chat"`
	Status	string	`json:"status" bson:"status"`
	Created time.Time `json:"created_at" bson:"created_at"`
	Updated time.Time `json:"updated_at" bson:"updated_at"`
}

type ChatUsers struct {
	Chat    *Chat `bson:"chat" json:"chat"`
	ChatID    primitive.ObjectID `bson:"chatId" json:"chatId"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	User    *User `bson:"user" json:"user"`
	IsAdmin bool `bson:"isAdmin" json:"isAdmin"`
}
