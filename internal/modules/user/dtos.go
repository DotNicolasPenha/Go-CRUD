package user

type CreateUserDTO struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Bio          string `json:"bio"`
}
type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
