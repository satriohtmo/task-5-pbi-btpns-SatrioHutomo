package app

type Register struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}