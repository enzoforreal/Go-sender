package services

import (
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

type MyEMAILProvider struct{}

func (e *MyEMAILProvider) SendEMAIL(from, to, object, body string) error {
	// implementation of SendEMAIL Method using a third-party EMAIL provider API
	//use twilio another api email

	return nil
}

func (e *MyEMAILProvider) SendMultipleEMAIL(from string, receivers []string, object, message string) error {
	//implementation of SendMultipleEMAIL method
	for _, to := range receivers {
		err := e.SendEMAIL(from, to, object, message)
		if err != nil {
			return err
		}
	}
	return nil
}

type TwilioSMSProvider struct {
	AccountSid string
	AuthToken  string
	FromNumber string
}

type MyCreateMessageParams struct {
	Body *MyCreateMessageParamsBody `json:"body"`
	From string                     `json:"from"`
	To   string                     `json:"to"`
}

type MyCreateMessageParamsBody struct {
	Message string `json:"message"`
}

func NewTwilioSMSProvider() *TwilioSMSProvider {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_FROM_NUMBER")

	return &TwilioSMSProvider{
		AccountSid: accountSid,
		AuthToken:  authToken,
		FromNumber: fromNumber,
	}
}

func (p *TwilioSMSProvider) SendSMS(phoneNumber string, message string) error {
	// Implementation of SendSMS method using a third-party SMS provider API
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: p.AccountSid,
		Password:   p.AuthToken,
	})
	msg := message
	params := &MyCreateMessageParams{
		Body: &MyCreateMessageParamsBody{
			Message: msg,
		},
		From: p.FromNumber,
		To:   phoneNumber,
	}

	createParams := &twilioApi.CreateMessageParams{
		Body: &msg,
		From: &params.From,
		To:   &params.To,
	}

	_, err := client.Api.CreateMessage(createParams)
	if err != nil {
		log.Fatal("Error sending SMS: ", err.Error())
		return err
	}

	return nil
}

func (p *TwilioSMSProvider) SendMultipleSMS(phoneNumbers []string, message string) error {
	//implementation of SendMultipleSMS method
	for _, phoneNumber := range phoneNumbers {
		err := p.SendSMS(phoneNumber, message)
		if err != nil {
			return err
		}
	}

	return nil
}
