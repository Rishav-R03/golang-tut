package main

import "fmt"

type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct {
	email string
}

type SmsNotifier struct {
	phone string
}

func (e *EmailNotifier) Send(msg string) error {
	if e.email == "" || e == nil {
		return fmt.Errorf("email cannot be empty")
	}
	fmt.Println("email sent with msg: ", msg)
	return nil
}

func (p *SmsNotifier) Send(msg string) error {
	if p.phone == "" || p == nil {
		return fmt.Errorf("phone cannot be empty")
	}
	fmt.Println("sms sent with msg", msg)
	return nil
}

func NotifyUser(n Notifier, msg string) {
	err := n.Send(msg)
	if err != nil {
		fmt.Printf("failed to send: %v\n", err)
		return
	}
	fmt.Println("Notification status: Success")
}
func main() {
	// Initialize our concrete implementations
	myEmail := &EmailNotifier{email: "dev@example.com"}
	mySms := &SmsNotifier{phone: "+123456789"}

	fmt.Println("--- Individual Notifications ---")
	// We can pass different types to the same function
	NotifyUser(myEmail, "Hello via Email!")
	NotifyUser(mySms, "Hello via SMS!")

	fmt.Println("\n--- Bulk Notifications (The Power of Interfaces) ---")
	// We can store different types in a single slice of the interface
	services := []Notifier{
		myEmail,
		mySms,
		&EmailNotifier{email: "boss@company.com"},
	}

	for _, service := range services {
		service.Send("System Update: Version 2.0 is live!")
	}

	fmt.Println("\n--- Error Handling Test ---")
	badSms := &SmsNotifier{phone: ""}
	NotifyUser(badSms, "This will fail")
}
