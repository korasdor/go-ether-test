package models

type SignUpData struct {
	Username string `json:"username"`
	Login    string `json:"logun"`
	Password string `json:"password"`
}

type SignInData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Login    string `json:"login"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
	BindingToken string
}
