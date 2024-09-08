package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"

	"github.com/gorilla/websocket"
)

const totalConnections = 3

func StartConnection(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log := log.New(os.Stdout, fmt.Sprintf("[conn #%d] ", id), log.Lmsgprefix)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "goapp/ws",
	}

	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				return
			}
			log.Printf("%s", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}

			<-done
			return
		}
	}
}

func main() {

	fmt.Println("Started WS client")

	// Wait for all connections to finish
	var wg sync.WaitGroup

	for i := 1; i <= totalConnections; i++ {
		wg.Add(1)
		go StartConnection(i, &wg)
	}

	wg.Wait()
}
