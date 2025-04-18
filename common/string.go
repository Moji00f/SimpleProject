package common

import (
	"github.com/Moji00f/SimpleProject/config"
	"golang.org/x/exp/rand"
	"math"
	"strconv"
	"time"
)

func GenerateOtp() string {
	cfg := config.GetConfig()
	rand.Seed(uint64(time.Now().UnixNano()))
	min := int(math.Pow10(cfg.Otp.Digits - 1))
	max := int(math.Pow10(cfg.Otp.Digits) - 1)

	var num = rand.Intn(max-min) + min

	return strconv.Itoa(num)
}
