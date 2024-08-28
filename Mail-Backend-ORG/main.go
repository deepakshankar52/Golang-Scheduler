package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
)

type Student struct {
	Name       string
	LeaveCount int
	Email      string
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		fmt.Println("Error retrieving the file:", err)
		return
	}
	defer file.Close()

	// Save the uploaded file
	tempFile, err := os.Create(filepath.Join("./uploads", handler.Filename))
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		fmt.Println("Error saving the file:", err)
		return
	}
	defer tempFile.Close()

	io.Copy(tempFile, file)
	fmt.Println("File uploaded successfully:", handler.Filename)

	// Process the file
	processExcelFile(filepath.Join("./uploads", handler.Filename), "./mail_template.html")
}

func processExcelFile(filePath string, templatePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error reading Excel rows:", err)
		return
	}

	for _, row := range rows[1:] { // Skipping the header row
		name := row[0]
		leaveCount, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("Invalid leave count for", name)
			continue
		}
		email := row[2]

		if leaveCount > 5 {
			student := Student{Name: name, LeaveCount: leaveCount, Email: email}
			sendGomail(student, templatePath)
		}
	}
}

func sendGomail(student Student, templatePath string) {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	t.Execute(&body, struct {
		Name       string
		LeaveCount int
	}{
		Name:       student.Name,
		LeaveCount: student.LeaveCount,
	})

	m := gomail.NewMessage()
	m.SetHeader("From", "deepaksciencee@gmail.com")
	m.SetHeader("To", student.Email)
	m.SetHeader("Subject", "Leave Warning Notification")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "deepaksciencee@gmail.com", "<dummy password>")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email to", student.Name, ":", err)
	} else {
		fmt.Println("Email sent to", student.Name)
	}
}

func main() {
	http.HandleFunc("/api/upload", uploadFileHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
