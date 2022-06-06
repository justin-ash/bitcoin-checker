package Utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

/*
	trigger current date with time
*/
func getCurrentTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

/*
	get any environment variable with key
*/
func getEnvVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(".env file not found")
	}

	return os.Getenv(key)
}

/*
	for sending mails
*/
func sendMail(message, toEmail string) {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", getEnvVariables("MAIL_USERNAME"), getEnvVariables("MAIL_PASSWORD"), getEnvVariables("MAIL_HOST"))

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{toEmail}
	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: Cryptocurrency Price Tracker\r\n" +
		"Hi User, \r\n There is an update for your bitcoin \r\n" +
		message)
	err := smtp.SendMail(getEnvVariables("MAIL_HOST")+":"+getEnvVariables("MAIL_PORT"), auth, getEnvVariables("MAIL_USERNAME"), to, msg)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
}

/*
	convert string to int
*/
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
