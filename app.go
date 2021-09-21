package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"goDevs/parser"
	"log"
	"net/http"
)

func HelloView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go Devs!")
}

func ParserView(w http.ResponseWriter, r *http.Request) {
	data := parser.Parser("godevs")
	json.NewEncoder(w).Encode(data)
}

func main() {
	go func() {
		GetAllPersonEverySec(30)
	}()

	handler := mux.NewRouter()
	handler.HandleFunc("/", HelloView).Methods("GET")
	handler.HandleFunc("/list", ParserView).Methods("GET")

	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		log.Fatalln("Listen and Serve problem..")
	}
}

func GetAllPersonEverySec(second uint64) {
	cron := gocron.NewScheduler()
	err := cron.Every(second).Seconds().Do(parser.Parser, "godevs")
	if err != nil {
		fmt.Println("GoCron Problem")
	}
	<-cron.Start()
}
