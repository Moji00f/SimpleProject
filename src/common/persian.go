package common

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/Moji00f/SimpleProject/config"
	"github.com/Moji00f/SimpleProject/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

const iranianMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidator(mobileNumber string) bool {

	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}

	return res
}

func GenerateOpt() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))   //10^d-1 --> 100000
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1) //999999 = 1000000 - 1 --> (10^d)-1

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
