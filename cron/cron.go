package cron

import (
	"MyTransactAPP/config"
	"MyTransactAPP/utils"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func generateExcelReport(db *gorm.DB) (string, error) {
	// Logic to generate Excel report
	return utils.GenerateExcelReport(db)
}

func dailyReportTask(db *gorm.DB) {
	// Generate Excel report
	filePath, err := generateExcelReport(db)
	if err != nil {
		config.Log.Errorf("Error generating Excel report: %v\n", err)
		return
	}

	// Send email with the report
	recipient := "user@example.com" // replace with the actual recipient's email address
	err = utils.SendTransactionEmail(filePath, recipient)
	if err != nil {
		config.Log.Errorf("Error sending email: %v\n", err)
		return
	}

	config.Log.Info("Daily transaction report sent successfully")
}

func SetupCron(db *gorm.DB) {
	c := cron.New()
	config.Log.Info("The cron setup function___________________________________")
	_, err := c.AddFunc("0 0 * * *", func() { dailyReportTask(db) }) // Runs at midnight every day
	if err != nil {
		config.Log.Errorf("Error scheduling the cron job: %v\n", err)
	}
	c.Start()
}
