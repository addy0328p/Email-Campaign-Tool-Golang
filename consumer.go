package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "भइया, चिंता मत करा। ई सब बिल्कुल नॉर्मल बा। जब नया भाषा (Go) सीखे लगेला, ऊपर से Docker, SMTP, CSV Parsing सब एक साथ करे के कोशिश होला, त हालत थोड़ा खराब होखे लागेला।")
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker :%d Error parsing template for %s", id, recipient.Email)
			// todo: add to dlq
			continue
		}

		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)

		err = smtp.SendMail(
			smtpHost+":"+smtpPort,
			nil,
			"ourangseb2@gmail.com",
			[]string{recipient.Email},
			[]byte(msg),
		)

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Sent email to %s \n", id, recipient.Email)

	}
}
