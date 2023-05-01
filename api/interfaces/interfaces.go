package interfaces

type SMSProvider interface {
	SendSMS(phoneNumber, messageContent string /*,name string*/) error
	SendMultipleSMS(phoneNumbers []string, messageContent string /*,name string*/) error
}

type EMAILProvider interface {
	SendEMAIL(from, to, object, messageContent string) error
	SendMultipleEMAIL(from string, receivers []string, object, messageContent string) error
}
