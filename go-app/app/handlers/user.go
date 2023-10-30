package handlers

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"

	"projector-test-app/internal/models"
	"projector-test-app/internal/repository"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (h *User) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	gofakeit.Struct(&newUser)
	newUser.CreatedAt = time.Now()

	userRepo := repository.NewUser(h.db)
	if _, err := userRepo.Save(r.Context(), &newUser); err != nil {
		panic(err)
	}

	w.Write([]byte("Success - OK\n"))
}
