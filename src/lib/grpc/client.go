package grpc

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client - gRPC client
type Client struct {
	Addr string
}

// NewClient - setups the client
func NewClient(host string, port string) *Client {
	addr := strings.Join([]string{host, ":", port}, "")
	return &Client{addr}
}

func (c *Client) dial() *grpc.ClientConn {
	conn, err := grpc.Dial(c.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	}

	log.WithFields(log.Fields{
		"result": "Connected!",
	}).Debug("gRPC")

	return conn
}
