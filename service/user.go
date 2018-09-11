package service

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}
