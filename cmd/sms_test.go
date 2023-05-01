package cmd

import (
	"app-admin-sender/api/services"
	"fmt"
	"log"
	"testing"
)

func TestSendSMS(t *testing.T) {
	smsProvider := services.NewTwilioSMSProvider()

	phoneNumber := "+33758589457" // num du destinataire
	message := "Test message"     // message a envoyer

	err := smsProvider.SendSMS(phoneNumber, message)
	if err != nil {
		log.Fatal("Error sending SMS:", err)
	}

	fmt.Println("SMS sent successfully")

}
