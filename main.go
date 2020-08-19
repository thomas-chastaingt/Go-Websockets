package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("hello world")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Println("new connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("Message receive from the client" + msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
		return nil
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
