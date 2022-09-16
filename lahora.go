package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beevik/ntp"
)

var ntp_server = "time.eljojo.net"

func hello(w http.ResponseWriter, req *http.Request) {
	response, _ := ntp.Query(ntp_server)
	time := time.Now().Add(response.ClockOffset).Format(time.RFC3339Nano)
	fmt.Fprintf(w, "%s\n%v\nprecision: %v", ntp_server, time, response.Precision)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
