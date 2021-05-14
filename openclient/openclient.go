package openclient

import (
	"bufio"
	"crypto/tls"
	"ctrader/stubs/model"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	Address      string
	ClientID     string
	ClientSecret string
	CertFile     string
	KeyFile      string
	reader       *bufio.Reader
	connected    bool
	conn         *tls.Conn
}

func (c *Client) Connect() error {
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		log.Fatalf("Unable to load cert or key file: %s", err)
		return err
	}
	c.conn, err = tls.Dial("tcp", c.Address, &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatalf("Unable to connect to host: %s", err)
		return err
	}
	c.connected = true
	c.reader = bufio.NewReader(c.conn)
	go c.sendHearBeats()
	return nil
}

func (c *Client) Disconnect() {
	c.conn.Close()
	c.connected = false
}

func (c *Client) SendMessage(m *model.ProtoMessage) error {
	data, _ := proto.Marshal(m)
	lenPacket := ToByteArray(proto.Size(m))
	Reverse(lenPacket)
	//@todo: Add error handling
	c.conn.Write(lenPacket)
	c.conn.Write(data)
	return nil
}

// Reads a single message
func (c *Client) ReadMessage() (*model.ProtoMessage, error) {
	// read message length
	firstFourBytes := make([]byte, 4)
	_, err := c.reader.Read(firstFourBytes)
	if err != nil {
		return nil, err
	}

	messageLen := 0
	for _, val := range firstFourBytes {
		messageLen += int(val)
	}
	// read message content
	packet := make([]byte, messageLen)
	_, err = c.reader.Read(packet)
	messagePb := &model.ProtoMessage{}
	err = proto.Unmarshal(packet, messagePb)
	if err != nil {
		return nil, err
	}
	return messagePb, nil
}

func (c *Client) sendHearBeats() {
	for range time.Tick(time.Second * 30) {
		if c.connected {
			message := &model.ProtoMessage{}
			event := &model.ProtoHeartbeatEvent{}
			payload, err := proto.Marshal(event)
			if err != nil {
				log.Println("Ping message marshal failed", err)
			}
			payloadType := uint32(model.ProtoPayloadType_HEARTBEAT_EVENT)
			message.Payload = payload
			message.PayloadType = &payloadType
			c.SendMessage(message)
		}
	}
}
