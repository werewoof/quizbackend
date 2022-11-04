package session

type Session struct {
	//	Expires int64  `json:"expires"`
	Token string `json:"token"`
	Id    int    `json:"id"`
}
