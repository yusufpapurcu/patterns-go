package writer

import "fmt"

type DisplayWriter struct{}

func (d *DisplayWriter) WriteEmails(emails ...string) error {
	for _, email := range emails {
		fmt.Println(email)
	}
	return nil
}

func (d *DisplayWriter) Close() error {
	return nil
}
