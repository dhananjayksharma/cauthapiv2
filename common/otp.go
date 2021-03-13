package common

import (
	"strconv"
	"time"
	"github.com/xlzd/gotp"
)

// func to generate otp
func GenerateOtp() string{
	secretLength := 16
	totp := gotp.NewDefaultTOTP(gotp.RandomSecret(secretLength))
	otp := totp.Now()
	return otp
}

// func to validate otp
func ValidateOtp(validity string) float64{
	otptime :=time.Now()
	i, err := strconv.ParseInt(validity, 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)
	// get the diff
	diff := tm.Sub(otptime)
	timediff := diff.Minutes()
	return timediff
}