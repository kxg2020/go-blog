package service

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

type Condition struct {
	Username     string `json:"username,omitempty"`
	Status       string `json:"status,omitempty"`
	Create_time  []int  `json:"date,omitempty"`
}

type Search struct {
	Condition    Condition
	Page         int    `json:"page"`
	Size         int    `json:"size"`
}
