package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func sendSchedulerMailHTML(subject string, templatePath string, to []string) {
	// Get HTML Data
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Deepak"})
	if err != nil {
		fmt.Println("Error(mSSM001): ", err)
		return
	}

	auth := smtp.PlainAuth(
		"",                         // Identity
		"deepaksciencee@gmail.com", // Username
		"emsltlshwapraowk",         // Scheduler - app password
		"smtp.gmail.com",           // Host for gmail
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	lerr := smtp.SendMail(
		"smtp.gmail.com:587", // Address of smtp server
		auth,
		"deepaksciencee@gmail.com", // From address
		to,                         // To Slice
		[]byte(msg),
	)

	if lerr != nil {
		fmt.Println("Error(mSSM002): ", lerr)
	}
}

func sendGomail(templatePath string) {
	// Get HTML Data
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Deepak"})
	if err != nil {
		fmt.Println("Error(mSSM001): ", err)
		return
	}

	// Send mail using gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "deepaksciencee@gmail.com")
	m.SetHeader("To", "deepakshankar52@gmail.com")
	// m.SetHeader("To", "kabiland04@gmail.com")
	m.SetHeader("Subject", "Hello Buddy!")
	m.SetBody("text/html", body.String())
	// m.Attach("./Assets/Sujatha_Img.png")

	d := gomail.NewDialer("smtp.gmail.com", 587, "deepaksciencee@gmail.com", "emsltlshwapraowk")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func main() {
	// sendSchedulerMail(
	// 	"Second mail buddy",
	// 	"Content for second mail body",
	// 	[]string{"deepakshankar52@gmail.com"}
	// )

	// sendSchedulerMailHTML(
	// 	"Fourth Mail Buddy",
	// 	"./test.html",
	// 	[]string{"deepakshankar52@gmail.com"},
	// )

	// sendGomail("./test.html")

	sendGomail("./mail_template.html")

	// fmt.Println("hello world")
}
