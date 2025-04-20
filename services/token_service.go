package services

import (
	"github.com/Moji00f/SimpleProject/api/dto"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constant"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"github.com/Moji00f/SimpleProject/pkg/service_errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: logger,
	}
}

func (t *TokenService) GenerateToken(token tokenDto) (*dto.TokenDetail, error) {
	tokenDetail := &dto.TokenDetail{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(t.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(t.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaim := jwt.MapClaims{}
	accessTokenClaim[constant.UserIdKey] = token.UserId
	accessTokenClaim[constant.FirstNameKey] = token.FirstName
	accessTokenClaim[constant.LastNameKey] = token.LastName
	accessTokenClaim[constant.UsernameKey] = token.Username
	accessTokenClaim[constant.EmailKey] = token.Email
	accessTokenClaim[constant.MobileNumberKey] = token.MobileNumber
	accessTokenClaim[constant.RolesKey] = token.Roles
	accessTokenClaim[constant.ExpireTimeKey] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)
	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(t.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaim := jwt.MapClaims{}
	refreshTokenClaim[constant.UserIdKey] = token.UserId
	refreshTokenClaim[constant.ExpireTimeKey] = tokenDetail.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)
	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(t.cfg.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func (t *TokenService) VerifyToken(token string) (*jwt.Token, error) {

	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(t.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (t *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {

	claimMap = map[string]interface{}{}

	verifyToke, err := t.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToke.Claims.(jwt.MapClaims)
	if ok && verifyToke.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}

		return claimMap, nil
	}

	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
