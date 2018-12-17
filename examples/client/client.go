package main

import (
	"fmt"
	"net/rpc"
	"os"
)

//SendMailArg argument to Mail.Send function
type SendMailArg struct {
	Recepient string
	Body      []byte
	Login     string
	Password  string
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err == nil {
		var result int
		arg := SendMailArg{
			Body:      []byte("simple example"),
			Recepient: os.Getenv("RECEPIENT_EMAIL"),
			Login:     os.Getenv("SENDER_EMAIL"),
			Password:  os.Getenv("SENDER_PASSWORD"),
		}
		err = client.Call("Mail.Send", arg, &result)

		if err != nil {
			fmt.Printf("%v", err)
			panic("error ")
		}
		fmt.Printf("Result is %d", result)
	} else {
		fmt.Printf("Error during rpc.Dial execution: %v ", err)
	}

}
