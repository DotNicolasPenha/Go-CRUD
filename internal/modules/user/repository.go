package user

import (
	"context"
	"time"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	if conn == nil {
		logger.Fatal("connection of repository user is nil")
	}
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Insert(createUserDto CreateUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.conn.Exec(
		ctx, "INSERT INTO users (nameuser,bio) VALUES ($1,$2)", createUserDto.Username, createUserDto.Bio,
	)
	return err
}
