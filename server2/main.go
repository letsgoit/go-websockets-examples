package main

import (
	"encoding/csv"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func getCSVReader() *csv.Reader{
	f, err := os.Open("./data/snp-500-intraday-data/dataset.csv")
	if err != nil {
		log.Print("loadCSV: ", err)
		return nil
	}

	r := csv.NewReader(f)
	return r
}

func sendMarketData(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	c, err:= upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()

	csvReader := getCSVReader()
	for {
		record, err := csvReader.Read()
		if err != nil {
			log.Print("readCSV: ", err)
			return
		}

		if len(record) > 0 {
			err = c.WriteMessage(websocket.TextMessage, []byte(record[0]+ "," + strings.Join(record[6:11], ",")))
			if err != nil {
				log.Print("write: ", err)
				return
			}
			time.Sleep(time.Second)
		}
	}

}

func main() {
	http.HandleFunc("/", sendMarketData)
	http.ListenAndServe(*addr, nil)
}