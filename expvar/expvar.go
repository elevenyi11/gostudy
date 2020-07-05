package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

var visits = expvar.NewInt("visits")
var stats = expvar.NewMap("http")
var requests, requestsFailed expvar.Int

func init() {
	stats.Set("req_succ", &requests)
	stats.Set("req_failed", &requestsFailed)
}

var start = time.Now()

func calculateUptime() interface{} {
	return time.Since(start).String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, i love %s!", r.URL.Path[1:])
}

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", mux)
	expvar.Publish("uptime", expvar.Func(calculateUptime))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
