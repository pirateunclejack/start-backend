package model

type SignupRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
