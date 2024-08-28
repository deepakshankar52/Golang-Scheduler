package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
)

type Student struct {
	Name   string
	Leaves int
	Email  string // Assuming you have emails in the excel file too
}

func sendGomail(student Student, templatePath string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error(mSSM001): ", err)
		return
	}

	// Execute the template with student data
	t.Execute(&body,
		struct {
			Name       string
			LeaveCount int
		}{
			Name:       student.Name,
			LeaveCount: student.Leaves,
		})

	m := gomail.NewMessage()
	m.SetHeader("From", "deepaksciencee@gmail.com")
	m.SetHeader("To", student.Email)
	m.SetHeader("Subject", "Leave Alert: Action Required")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "deepaksciencee@gmail.com", "emsltlshwapraowk")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email to", student.Name, ":", err)
	} else {
		fmt.Println("Email sent to", student.Name)
	}
}

func readExcelAndSendEmails(filePath string, templatePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Assuming the sheet is named "Sheet1"
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	for _, row := range rows[1:] { // Skipping the header row
		name := row[0]
		leaveCount := row[1] // Assuming this is a string representation of the leave count
		email := row[2]      // Assuming the email is in the third column

		leaveCountInt, err := strconv.Atoi(leaveCount)
		if err != nil {
			fmt.Println("Invalid leave count for", name)
			continue
		}

		if leaveCountInt == 5 {
			student := Student{Name: name, Leaves: leaveCountInt, Email: email}
			sendGomail(student, templatePath)
		}
	}
}

func main() {
	readExcelAndSendEmails("./Assets/Student_Attendance.xlsx", "./mail_template.html")
}
