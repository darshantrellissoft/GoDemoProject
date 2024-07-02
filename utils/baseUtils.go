package utils

import (
	"MyTransactAPP/config"
	"MyTransactAPP/models"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type EmailConfig struct {
	From     string
	SMTPHost string
	SMTPUser string
	SMTPPass string
}

var EmailSettings EmailConfig

func GetEmailConfig() EmailConfig {
	EmailSettings = EmailConfig{
		From:     os.Getenv("EMAIL_FROM"),
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPass: os.Getenv("SMTP_PASS"),
	}
	if EmailSettings.From == "" || EmailSettings.SMTPHost == "" || EmailSettings.SMTPUser == "" || EmailSettings.SMTPPass == "" {
		config.Log.Error("Email configuration variables missing")
	}
	return EmailSettings
}

func SendWelcomeEmail(user models.User) error {
	tmpl, err := template.ParseFiles("templates/welcome.html")
	if err != nil {
		return err
	}

	var body strings.Builder
	if err := tmpl.Execute(&body, user); err != nil {
		return err
	}

	return sendEmail(user.Email, "Welcome to MyTransactAPP", body.String())
}

func sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	emailConfig := GetEmailConfig()
	m.SetHeader("From", emailConfig.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		emailConfig.SMTPHost,
		587,
		emailConfig.SMTPUser,
		emailConfig.SMTPPass,
	)

	return d.DialAndSend(m)
}

func SendConfirmationEmail(email, link string) error {
	tmpl, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		return err
	}

	var body strings.Builder
	if err := tmpl.Execute(&body, map[string]string{"Link": link}); err != nil {
		return err
	}

	return sendEmail(email, "Payment Confirmation", body.String())
}

type Transaction struct {
	ID         string
	CreatedAt  time.Time
	CardNumber string
	ExpiryDate string
	CVV        string
	Amount     float64
	Status     string
	UserID     uint
	// Add other fields if needed
}

func GenerateExcelReport(db *gorm.DB) (string, error) {
	f := excelize.NewFile()
	sheetName := "Transactions"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return "", fmt.Errorf("failed to create new sheet: %w", err)
	}
	f.SetActiveSheet(index)

	// Set headers
	headers := []string{"Transaction ID", "Created At", "Card Number", "Expiry Date", "Amount", "Status"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheetName, cell, header)
	}

	// Fetch transactions from database
	var transactions []Transaction
	result := db.Find(&transactions)
	if result.Error != nil {
		return "", result.Error
	}

	// Add transactions data to the sheet
	for i, transaction := range transactions {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), transaction.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), transaction.CreatedAt.Format(time.RFC3339))
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), transaction.CardNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), transaction.ExpiryDate)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), transaction.Amount)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), transaction.Status)
	}

	// Save the file
	filePath := "transactions_report.xlsx"
	if err := f.SaveAs(filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

func SendTransactionEmail(filePath string, recipient string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "Daily Transaction Report")
	m.SetBody("text/plain", "Please find attached the daily transaction report.")
	m.Attach(filePath)

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))

	return d.DialAndSend(m)
}
