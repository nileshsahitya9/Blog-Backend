package helper

import (
	"net/smtp"
	"os"
)

func SendOTP(email, otp string) error {

	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := "From: " + from + "\n" +
		"To: " + email + "\n" +
		"Subject: One-Time Password (OTP)\n\n" +
		"Thank you for signing into the Strugglers Blog. We sincerely hope you have a pleasant experience during your time here. As part of the authentication process, we have generated a One-Time Password (OTP) for you. Please use the following OTP to proceed: " + otp

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
