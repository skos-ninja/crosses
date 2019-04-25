package user

import (
	"context"
	"time"

	"crosses/db"
	crossErr "crosses/err"
	"crosses/services/user/models"

	"github.com/cuvva/ksuid"
)

type Service interface {
	GetUser(ctx context.Context, userID string) (*models.Player, error)
	CreateUser(ctx context.Context, name string) (*models.Player, error)
}

type userService struct {
	db *db.DB
}

func NewUserService(db *db.DB) Service {
	return &userService{db: db}
}

func (svc *userService) GetUser(ctx context.Context, userID string) (*models.Player, error) {
	row := svc.db.QueryRowWithCtx(ctx, "SELECT id, name, created_at FROM players WHERE id=$1", userID)
	if row != nil {
		user := &models.Player{}

		err := row.Scan(&user.ID, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, crossErr.New("user_not_found", nil)
}

func (svc *userService) CreateUser(ctx context.Context, name string) (*models.Player, error) {
	id := ksuid.Generate("player")
	now := time.Now().UTC()
	user := models.Player{
		ID:        id.String(),
		Name:      name,
		CreatedAt: &now,
	}

	_, err := svc.db.Exec(ctx, "INSERT INTO players (id, name, created_at) VALUES ($1, $2, $3)", user.ID, user.Name, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
