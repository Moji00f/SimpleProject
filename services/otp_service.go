package services

import (
	"errors"
	"fmt"
	"github.com/Moji00f/SimpleProject/common"
	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/constant"
	"github.com/Moji00f/SimpleProject/data/cache"
	"github.com/Moji00f/SimpleProject/pkg/logging"
	"github.com/Moji00f/SimpleProject/pkg/service_errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type otpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	getRedis := cache.GetRedis()

	return &OtpService{logger: logger, cfg: cfg, redisClient: getRedis}
}

func (o *OtpService) SendOtp(mobileNumber string) error {
	otp := common.GenerateOtp()
	err := o.SetOtp(mobileNumber, otp)
	if err != nil {
		return err
	}

	return nil
}

func (o *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constant.RedisOtpDefaultKey, mobileNumber)
	val := &otpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[otpDto](o.redisClient, key)
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	if err == nil && res.Value != "" && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OptExists}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}
	if res.Value == "" && errors.Is(err, redis.Nil) {

		err = cache.Set(o.redisClient, key, val, o.cfg.Otp.ExpireTime*time.Second)
		return err
	}
	return nil
}

func (o *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constant.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[otpDto](o.redisClient, key)
	if err != nil {
		return err
	} else if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(o.redisClient, key, res, o.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
