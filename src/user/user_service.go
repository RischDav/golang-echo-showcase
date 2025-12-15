package user

import (
	"context"
	"golang-echo-showcase/src/user/sqlc/output"
)

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) CreateUser(ctx context.Context, firstname, lastname string) (sqlc.User, error) {
	params := sqlc.CreateUserParams{
		Firstname: firstname,
		Lastname:  lastname,
	}
	return s.queries.CreateUser(ctx, params)
}

func (s *Service) GetUser(ctx context.Context, id int64) (sqlc.User, error) {
	return s.queries.GetUser(ctx, id)
}

func (s *Service) UpdateUserFirstname(ctx context.Context, id int64, firstname string) error {
	params := sqlc.UpdateUserFirstnameParams{
		Firstname: firstname,
		ID:        id,
	}
	return s.queries.UpdateUserFirstname(ctx, params)
}

func (s *Service) UpdateUserLastname(ctx context.Context, id int64, lastname string) error {
	params := sqlc.UpdateUserLastnameParams{
		Lastname: lastname,
		ID:       id,
	}
	return s.queries.UpdateUserLastname(ctx, params)
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.queries.DeleteUser(ctx, id)
}