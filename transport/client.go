package transport

import (
	"net"

	"github.com/isundaylee/flutterplot/data"
	"github.com/isundaylee/flutterplot/messenger"
)

func Start(addr *net.TCPAddr) {
	messenger, err := messenger.Make(addr)
	if err != nil {
		panic(err)
	}

	messenger.AddHandler("data", func(rawEntries interface{}) {
		entries := rawEntries.([]interface{})

		for _, rawEntry := range entries {
			entry := rawEntry.(map[string]interface{})

			_ = entry["entity"].(string)
			metric := entry["metric"].(string)
			value := entry["value"].(float64)

			if metric == "rx_bytes" {
				data.AddDataPoint(value)
			}
		}
	})

	messenger.Start()
}
