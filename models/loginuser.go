package models

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDetails struct {
	PkUserweb int    `json:"pk_userweb"`
	Username  string `json:"username"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
}
