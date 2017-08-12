package transport

import (
	"encoding/binary"
	"errors"
	"net"

	"github.com/isundaylee/flutterplot/data"
	"github.com/isundaylee/flutterplot/flutter"

	"github.com/golang/protobuf/proto"
)

func Connect(addr *net.TCPAddr) {
	go func() {
		client, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			panic(err)
		}

		for {
			var buf = make([]byte, 1024)
			n, err := client.Read(buf)
			if err != nil {
				panic(err)
			}

			length := binary.BigEndian.Uint32(buf)
			if length+4 != uint32(n) {
				panic(errors.New("Unmatching flutter message length"))
			}

			protoData := &flutter.Data{}
			if err := proto.Unmarshal(buf[4:length+4], protoData); err != nil {
				panic(err)
			}

			for _, entry := range protoData.GetPoints() {
				if entry.GetMetric() == "tx_bytes" {
					data.AddDataPoint(entry.GetValue())
				}
			}
		}
	}()
}
