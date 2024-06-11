package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type LoginRes struct {
	Correct   bool `json:"correct"`
	InCorrect bool `json:"incorrect"`
}
