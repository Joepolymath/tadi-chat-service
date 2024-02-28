package chats

type GroupChatRequest struct {
    ChatName   string   `json:"chatName"`
    Users      []string `json:"users"`
    GroupAdmin string   `json:"groupAdmin"`
	 IsGroupChat bool		`json:"isGroupChat"`
}