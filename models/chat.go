package models

import "time"

// Chat instance
type Chat struct {
	ChatName      string    `bson:"chatName" json:"chatName"`
	IsGroupChat   bool      `bson:"isGroupChat" json:"isGroupChat"`
	Users         []User  `bson:"users" json:"users"`
	GroupAdmin    string    `bson:"groupAdmin" json:"groupAdmin"`
	Timestamps    time.Time `bson:"timestamps" json:"timestamps"`
}

// Message in each chat
type Message struct {
    Sender  string        `json:"sender" bson:"sender"`
    Content string        `json:"content" bson:"content"`
    Chat    Chat        `json:"chat" bson:"chat"`
    Created time.Time     `json:"created_at" bson:"created_at"`
    Updated time.Time     `json:"updated_at" bson:"updated_at"`
}