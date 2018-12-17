package smtpclient

import (
	"crypto/tls"
	"net/smtp"
)

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return (s.host + ":" + s.port)
}

func SendEmail(sender string, password string, recepient string, body []byte) error {
	// Connect to the remote SMTP server.
	smtpServer := SmtpServer{host: "smtp.mail.ru", port: "465"}
	auth := smtp.PlainAuth("", sender, password, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}
	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = client.Mail(sender); err != nil {
		return err
	}

	if err = client.Rcpt(recepient); err != nil {
		return err
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(body)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	client.Quit()
	return nil
}
