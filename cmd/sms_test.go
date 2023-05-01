package cmd

import (
	"app-admin-sender/api/services"
	"fmt"
	"log"
	"testing"
)

func TestSendMultipleSMS(t *testing.T) {
	smsProvider := services.NewTwilioSMSProvider()

	phoneNumbers := []string{
		"+33758589457",
		//"+33601249729",
	}
	url := "https://www.google.fr/"
	message := "Bonjour DEUXIEME TEST c'est enzo lol tkt C'est juste pour tester avec un lien  si tu reçois ca cest que moi aussi je l'ai reçu cest qu'on est good aussi haha, veuillez cliquer sur ce lien pour accéder à notre site : " + url // message a envoyer
	err := smsProvider.SendMultipleSMS(phoneNumbers, message)
	if err != nil {
		log.Fatal("Error sending SMS:", err)
	}

	fmt.Println("SMS sent successfully")

}
