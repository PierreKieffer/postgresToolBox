package main

import (
	// "bytes"
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"log"
	"time"
)

func waitForNotification(l *pq.Listener) {
	for {
		select {
		case n := <-l.Notify:
			log.Println("Received data from channel [", n.Channel, "] :")

			// extract payload from notification
			var payload map[string]interface{}
			json.Unmarshal([]byte(n.Extra), &payload)
			log.Println(payload)

			return
		case <-time.After(90 * time.Second):
			log.Println("Received no events for 90 seconds, checking connection")
			go func() {
				l.Ping()
			}()
			return
		}
	}
}

func main() {
	var conninfo string = "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"

	_, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Println("sql Open error")
		log.Println(err)
	}

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("report problem")
			log.Println(err.Error())
		}
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("events")
	if err != nil {
		log.Println(err)
	}

	log.Println("Start monitoring PostgreSQL...")
	for {
		waitForNotification(listener)
	}
}
