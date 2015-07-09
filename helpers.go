package hello

import (
	"github.com/derrickrahbar-wf/gostats/gostats"
	"net/http"
)

func helper(r *http.Request) {
	a := 4
	b := 4 + a
	client := gostats.NewStatsClient(r)
	client.Incr("test.foo.basr", int64(b))
	client.Close()
}
