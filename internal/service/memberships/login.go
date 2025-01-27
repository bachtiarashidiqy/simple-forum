package memberships

import (
	"context"
	"errors"
	"github.com/bachtiarashidiqy/simple-forum/internal/model/memberships"
	"github.com/bachtiarashidiqy/simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Login, req.Login)
	if err != nil {
		log.Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid username or email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Err(err).Msg("failed to create token")
		return "", err
	}
	return token, nil
}
