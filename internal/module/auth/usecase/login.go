package usecase

import (
	"context"
	"fmt"
	"oms-be/internal/module/auth/entity"
)

type LoginUseCase struct {
	repo ILoginRepo
}

func NewLoginUseCase(repo ILoginRepo) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Login(ctx context.Context, payload entity.LoginPayload) error {
	user, err := uc.repo.GetUserByEmail(ctx, payload.Email)

	fmt.Println("Handle logic business")
	fmt.Println(user)

	if err != nil {
		return err
	}

	// another logic business here

	return nil
}
