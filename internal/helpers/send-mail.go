package helpers

import "github.com/DapperBlondie/service-monitor/internal/channeldata"

// SendEmail sends an email
func SendEmail(mailMessage channeldata.MailData) {
	// if no sender specified, use defaults
	if mailMessage.FromAddress == "" {
		mailMessage.FromAddress = app.PreferenceMap["smtp_from_email"]
		mailMessage.FromName = app.PreferenceMap["smtp_from_name"]
	}

	job := channeldata.MailJob{MailMessage: mailMessage}
	app.MailQueue <- job
}
