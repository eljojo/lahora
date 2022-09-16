package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beevik/ntp"
)

var ntp_server = "192.168.90.56"

func hello(w http.ResponseWriter, req *http.Request) {
	response, _ := ntp.Query(ntp_server)
	time := time.Now().Add(response.ClockOffset).Format(time.RFC3339Nano)
	fmt.Fprintf(w, "time.eljojo.net\n%s\nNTP precision: %v\nNTP dispersion: %v", time, response.Precision, response.RootDispersion)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
