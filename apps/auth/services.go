package auth

import (
	"context"
	"online-shop/infra/response"
	"online-shop/internal/config"
)

type Repository interface {
	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)

	if err = authEntity.Validate(); err != nil {
		return
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encrytion.Salt)); err != nil {
		return
	}

	authModel, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if authModel.IsExist() {
		return response.ErrEmailAlreadyExist
	}

	err = s.repo.CreateAuth(ctx, authEntity)

	return
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {

	authEntity := NewFromLoginRequest(req)

	if err = authEntity.Validate(); err != nil {
		return
	}

	authModel, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		return
	}

	if err = authEntity.VerifyPasswordFromPlain(authModel.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = authModel.GenerateToken(config.Cfg.App.Encrytion.JWTSecret)

	return
}
