package nats

import "github.com/nats-io/nats.go"

type NATSClient struct {
	Conn *nats.Conn
	Js   nats.JetStreamContext
}

func ConnectNATS(url string) (*NATSClient, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &NATSClient{Conn: nc, Js: js}, nil
}

func (c *NATSClient) Publish(subject string, msg []byte) error {
	_, err := c.Js.Publish(subject, msg)
	return err
}

func (c *NATSClient) Subscribe(subject string, handler nats.MsgHandler) error {
	_, err := c.Js.Subscribe(subject, handler)
	return err
}
