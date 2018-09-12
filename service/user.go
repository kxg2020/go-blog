package service

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

type Search struct {
	Username     string `json:"username,omitempty"`
	Status       string `json:"status,omitempty"`
	Create_time  []int  `json:"date,omitempty"`
}
