package server

import (
	"net/http"

	"appengine"
	"github.com/derrickrahbar-wf/gostats/gostats"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	client := gostats.NewStatsClient(r)
	client_2 := gostats.NewStatsClient(r)

	client.Incr("test.c1.incr", 2)
	client_2.Incr("test.c2.incr", 3)

	client.Decr("test.c1.incr", -1)
	client_2.Decr("test.c2.incr", -1)

	client.Timing("test.c1.time", 4)
	client_2.Timing("test.c2.time", 10)

	client.Gauge("test.c1.gauge", 5)
	client_2.Gauge("test.c2.gauge", 5)

	client.Close()
	client_2.Close()

	c := appengine.NewContext(r)
	c.Infof("Sending Some stats")
}
