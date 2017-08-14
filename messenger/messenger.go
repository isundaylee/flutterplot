package messenger

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"time"
)

const (
	BufferSize         = 8192
	HeaderSize         = 4
	HeartBeatType      = "heartbeat"
	HeartBeatReplyType = "heartbeat_reply"

	HeartBeatInterval = 1 * time.Second
)

type Messenger struct {
	conn *net.TCPConn

	handlers map[string]Handler
}

type Handler func(interface{})

func Make(server *net.TCPAddr) (*Messenger, error) {
	messenger := Messenger{}

	messenger.handlers = make(map[string]Handler)

	var err error
	messenger.conn, err = net.DialTCP("tcp", nil, server)
	if err != nil {
		return nil, err
	}

	go messenger.doHeartbeat()

	return &messenger, nil
}

func (messenger *Messenger) Start() {
	go messenger.doReceive()
}

func (messenger *Messenger) Send(messageType string, content interface{}) {
	payload := map[string]interface{}{
		"type": messageType,
		"body": content,
	}

	json, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	payloadSize := len(json)
	headerBuffer := make([]byte, HeaderSize)
	binary.BigEndian.PutUint32(headerBuffer, uint32(payloadSize))

	messenger.writeBytes(headerBuffer, HeaderSize)
	messenger.writeBytes(json, payloadSize)
}

func (messenger *Messenger) doHeartbeat() {
	timer := time.NewTimer(0)

	for {
		<-timer.C
		messenger.Send(HeartBeatType, struct{}{})
		timer.Reset(HeartBeatInterval)
	}
}

func (messenger *Messenger) readBytes(buffer []byte, n int) {
	totalRead := 0
	for totalRead < n {
		read, err := messenger.conn.Read(buffer[totalRead:n])
		if err != nil {
			panic(err)
		}

		totalRead += read
	}
}

func (messenger *Messenger) writeBytes(buffer []byte, n int) {
	totalWritten := 0
	for totalWritten < n {
		written, err := messenger.conn.Write(buffer[totalWritten:n])
		if err != nil {
			panic(err)
		}

		totalWritten += written
	}
}

func (messenger *Messenger) doReceive() {
	headerBuffer := make([]byte, HeaderSize)
	dataBuffer := make([]byte, BufferSize)

	for {
		messenger.readBytes(headerBuffer, HeaderSize)
		payloadSize := int(binary.BigEndian.Uint32(headerBuffer))

		if payloadSize > BufferSize {
			panic("Messenger buffer overflow.")
		}

		messenger.readBytes(dataBuffer, payloadSize)

		message := make(map[string]interface{})
		json.Unmarshal(dataBuffer[:payloadSize], &message)
		messageType := message["type"].(string)

		if messageType == HeartBeatType {
			messenger.Send(HeartBeatReplyType, message["body"])
			continue
		} else if messageType == HeartBeatReplyType {
			continue
		}

		handler, ok := messenger.handlers[messageType]
		if !ok {
			panic("Messenger encountered unknown message type: " + messageType)
		}

		handler(message["body"])
	}
}

func (messenger *Messenger) AddHandler(messageType string, handler Handler) {
	messenger.handlers[messageType] = handler
}
