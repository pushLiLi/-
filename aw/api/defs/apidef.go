package defs

// request
type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// Data Model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

// comment
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

// session
type SimpleSession struct {
	Username string //登录的用户名
	TTL      int64  //判断是否已经拥有此用户
}

// response
type SignedUp struct {
	Success   bool   `json:"sucess"`
	SessionId string `json:"session_Id"`
}
