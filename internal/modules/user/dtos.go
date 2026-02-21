package user

type CreateUserDTO struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
}
