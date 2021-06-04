package config

import "google.golang.org/grpc"

type RPCConfig struct {
	*grpc.ClientConn
	Host string `mapstructure:"GRPC_HOST" yaml:"host" env:"GRPC_HOST"`
	Port string `mapstructure:"GRPC_PORT" yaml:"port" env:"GRPC_PORT"`
}

func (c *RPCConfig) Setup() {
	addr := c.Host + ":" + c.Port
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	c.ClientConn = conn
}
