package models

import "time"

type Chat struct {
	ChatName      string    `bson:"chatName" json:"chatName"`
	IsGroupChat   bool      `bson:"isGroupChat" json:"isGroupChat"`
	Users         []string  `bson:"users" json:"users"`
	LatestMessage string    `bson:"latestMessage" json:"latestMessage"`
	GroupAdmin    string    `bson:"groupAdmin" json:"groupAdmin"`
	Timestamps    time.Time `bson:"timestamps" json:"timestamps"`
}
