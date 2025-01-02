package memberships

import (
	"context"
	"errors"
	"github.com/bachtiarashidiqy/simple-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *service) SingUp(ctx context.Context, req memberships.SingUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("email or username already exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	model := memberships.UserModel{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}
	err = s.membershipRepo.CreateUser(ctx, &model)
	if err != nil {
		return err
	}
	return nil
}
