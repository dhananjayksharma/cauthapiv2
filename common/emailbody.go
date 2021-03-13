package common

import (
	"fmt"
	"github.com/spf13/viper"
)

func RegisterBody(otp string) string{
	env := viper.GetString("ENV_VAR")
	var url string
	if env == "stage" {
		url = viper.GetString("url.stageenv")
	} else if env == "laptopenv" {
		url = viper.GetString("url.laptopenv")
	}	
	RegEmailBody := "Hi Cinestaan User, Welcome to Cinestaan! We're very excited to have you on board. Please Use this link to verify mail \n"+url+"/activateaccount\n . OTP for Verification is "+otp+". Need help, or have questions? Just reply to this email, we'd love to help. Yours truly,Cinestaan"
	fmt.Println(RegEmailBody)
	return RegEmailBody
}

func ForgetPasswordBody(otp string) string{
	env := viper.GetString("ENV_VAR")
	var url string
	if env == "stage" {
		url = viper.GetString("url.stageenv")
	} else if env == "laptopenv" {
		url = viper.GetString("url.laptopenv")
	}	
	RegEmailBody := "Hi Cinestaan User, Welcome to Cinestaan! We're very excited to have you on board. \n"+url+"/resetpassword\n. Use One Time Password :\n"+otp+". Need help, or have questions? Just reply to this email, we'd love to help. Yours truly,Cinestaan"
	fmt.Println(RegEmailBody)
	return RegEmailBody
}