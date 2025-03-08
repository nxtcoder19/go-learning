package main

import (
	"fmt"
)

type Comms interface {
	getWhatsAppNotification() string
	getEmailNotification() string
}

type Impl1 struct {
	whatsAppMsg string
	emailMsg    string
}

func (i *Impl1) getWhatsAppNotification() string {
	return i.whatsAppMsg
}

func (i *Impl1) getEmailNotification() string {
	return i.emailMsg
}

func newImpl1(whatsAppMsg, emailMsg string) Comms {
	return &Impl1{
		whatsAppMsg: whatsAppMsg,
		emailMsg:    emailMsg,
	}
}

func main() {
	fmt.Println("Try programiz.pro")

	d := newImpl1("Message from WhatsApp", "Message from Email")

	fmt.Println("WhatsApp Notification:", d.getWhatsAppNotification())
	fmt.Println("Email Notification:", d.getEmailNotification())
}
