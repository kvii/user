package user

import (
	"context"

	"github.com/kvii/pkg/database"
)

type Service struct {
	DB *database.DB
}

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	return "Hello " + name, nil
}
