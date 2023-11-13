package service

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gopkg.in/gomail.v2"
)

func (s *service) SendMail(ctx context.Context, emailTo []string, subject, emailBody string, file *excelize.File) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", s.smtpConfig.Email)
	msg.SetHeader("To", formatRecipients(emailTo))
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", emailBody)

	port, err := strconv.Atoi(s.smtpConfig.Port)
	if err != nil {
		return err
	}

	tempFile, err := ioutil.TempFile("", "*.xlsx")
	if err != nil {
		return err
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	err = file.SaveAs(tempFile.Name())
	if err != nil {
		return err
	}

	msg.Attach(tempFile.Name(), gomail.Rename(filepath.Base(tempFile.Name())))

	dialer := gomail.NewDialer(s.smtpConfig.Host, port, s.smtpConfig.Email, s.smtpConfig.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = dialer.DialAndSend(msg)
	return err
}

func buildMsg(emailTo []string, subject, emailBody string) []byte {
	msg := []byte(
		"To: " + formatRecipients(emailTo) + "\r\n" +
			"Subject: " + subject + "\r\n" +
			emailBody + "\r\n")

	return msg
}

func formatRecipients(recipients []string) string {
	formattedRecipients := ""
	for _, recipient := range recipients {
		formattedRecipients += recipient + ", "
	}

	formattedRecipients = formattedRecipients[:len(formattedRecipients)-2]
	return formattedRecipients
}
