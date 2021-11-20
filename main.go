package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getPodInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, greetings from kubernetes!"+"\n")

	getHostname := os.Hostname
	hostname, err := getHostname()
	gotError(err)
	fmt.Fprintf(w, "Pod name: %s", hostname+"\n")

	addrs, err := net.InterfaceAddrs()
	gotError(err)

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintf(w, "Pod IP: %s", ipnet.IP.String()+"\n")
			}
		}
	}

	appVersion := "v0.2"
	fmt.Fprintf(w, "App version: %s", appVersion)

}

func gotError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	addr := ":8080"
	http.HandleFunc("/", getPodInfo)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
