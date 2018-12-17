package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/acrap/gopostman/smtpclient"
)

//SendMailArg argument to send emails
type SendMailArg struct {
	Recepient string
	Body      []byte
	Login     string
	Password  string
}

//Mail implements simple mail sending methods
type Mail int

//Send method performs sending a simple email message
func (t *Mail) Send(args SendMailArg, res *int) error {
	fmt.Print("Trying to send message...")
	err := smtpclient.SendEmail(args.Login, args.Password, args.Recepient, args.Body)
	if err == nil {
		fmt.Printf("Message to %s has been sent", args.Recepient)
	}
	return nil
}

func main() {
	task := new(Mail)
	// Publish the receivers methods
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		panic("can't resolve tcp")
	}

	// Listen to TPC connections on port 1234
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic("Listen tcp error")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}
