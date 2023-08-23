package services

import (
	"fmt"
	"time"

	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constants"
	"github.com/Moji00f/SimpleProject/data/cache"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"github.com/Moji00f/SimpleProject/pkg/service_error"
	"github.com/go-redis/redis/v7"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()

	return &OtpService{logger: logger, cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_error.ServiceError{EndUserMessage: service_error.OptExists}
	} else if err == nil && res.Used {
		return &service_error.ServiceError{EndUserMessage: service_error.OtpUsed}
	}

	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}

	return nil

}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err != nil {
		return err
	} else if err == nil && res.Used {
		return &service_error.ServiceError{EndUserMessage: service_error.OtpUsed}
	} else if err == nil && !res.Used && res.Value != otp {
		return &service_error.ServiceError{EndUserMessage: service_error.OtpNotValid}
	} else if err == nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
