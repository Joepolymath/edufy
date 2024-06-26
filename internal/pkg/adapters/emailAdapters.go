package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	var err error
// 	cfg, err = config.Load()
// 	if err != nil {
// 		log.Println("error loading config::: ", zap.Error(err))
// 	}
// }

type IEmailAdapter interface {
	SendEmail(to, subject, body string) error
	// SendAccountCreateEmail(ctx context.Context, firstName, lastName, email, password string) (bool, error)
	// SendAdmissionDecline(ctx context.Context, firstName, lastName, email, schoolName string) (bool, error)
	// SendApplicationDecline(ctx context.Context, firstName, lastName, email, schoolName string) (bool, error)
}

// SendGridEmailAdapter is an implementation of the EmailAdapter interface using SendGrid.
type SendGridEmailAdapter struct {
	APIKey     string
	Sender     string
	APIBaseURL string
}

// NewSendGridEmailAdapter creates a new instance of SendGridEmailAdapter.
func NewSendGridEmailAdapter() *SendGridEmailAdapter {
	return &SendGridEmailAdapter{
		APIKey:     cfg.SendGridEmailAPIKey,
		Sender:     cfg.EmailHostUser,
		APIBaseURL: "https://api.sendgrid.com/v3/mail/send",
	}
}

// SendEmail sends an email using the SendGrid API.
func (s *SendGridEmailAdapter) SendEmail(to, subject, body string) error {
	// Create JSON payload
	payload, err := json.Marshal(map[string]interface{}{
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]string{
					{"email": to},
				},
				"subject": subject,
			},
		},
		"from": map[string]string{"email": s.Sender},
		"content": []map[string]string{
			{"type": "text/plain", "value": body},
		},
	})
	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", s.APIBaseURL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	// Send HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for successful response (status code 2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Email sent successfully!")
		return nil
	}

	// If the response code is not in the 2xx range, print an error message
	var responseBody []byte
	resp.Body.Read(responseBody)
	log.Printf("SendGrid API Error: %s\n", responseBody)
	return fmt.Errorf("failed to send email, status code: %d", resp.StatusCode)
}

// type EmailAdapter struct {
// }

// func NewEmailAdapter() IEmailAdapter {
// 	return &EmailAdapter{}
// }

// // SendEmail /* This is used for sending email */
// func (e *EmailAdapter) SendEmail(to, subject, body string) error {

// 	auth := smtp.PlainAuth("", cfg.EmailHostUser, cfg.EmailHostPassword, cfg.EmailHost)
// 	addr := fmt.Sprintf("%s:%s", cfg.EmailHost, cfg.EmailPort)

// 	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

// 	err := smtp.SendMail(addr, auth, cfg.EmailHostUser, []string{to}, message)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (e *EmailAdapter) SendAccountCreateEmail(ctx context.Context, firstName, lastName, email, password string) (bool, error) {
// 	/* This is used to email the user after creating an account */

// 	err := e.SendEmail(email, "Account Created", "Hello "+firstName+" "+lastName+",\n\nYour account has been created successfully.\n\nYour password is: "+password+"\n\nRegards,\nLearnium Team")
// 	if err != nil {
// 		// log.Println("error sending creation email for staff or student:::", zap.Error(err))

// 		return false, err
// 	}
// 	return true, nil
// }

// func (e *EmailAdapter) SendAdmissionDecline(ctx context.Context, firstName, lastName, email, schoolName string) (bool, error) {
// 	/* This is used to email the user after creating an account */

// 	err := e.SendEmail(email, "Admission Declined", "Hello "+firstName+" "+lastName+",\n\nYour admission have been declined to join "+schoolName+"\n\nRegards,\nLearnium Team")
// 	if err != nil {
// 		// utils.Error(ctx, "error sending declined  admission", zap.Error(err))

// 		return false, err
// 	}
// 	return true, nil
// }

// func (e *EmailAdapter) SendApplicationDecline(ctx context.Context, firstName, lastName, email, schoolName string) (bool, error) {
// 	/* This is used to email the user after creating an account */

// 	err := e.SendEmail(email, "Application Declined", "Hello "+firstName+" "+lastName+",\n\nYour application have been declined to join "+schoolName+"\n\nRegards,\nLearnium Team")
// 	if err != nil {
// 		// utils.Error(ctx, "error sending declined  admission", zap.Error(err))

// 		return false, err
// 	}
// 	return true, nil
// }
