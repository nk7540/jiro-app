package config

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// Environment - environment variables
type Environment struct {
	Port              string `envconfig:"PORT" default:"8080"`
	MysqlUser         string `envconfig:"MYSQL_USER" default:"root"`
	MysqlPassword     string `envconfig:"MYSQL_PASSWORD" required:"true"`
	MysqlHost         string `envconfig:"MYSQL_HOST" default:"artics-db"`
	MysqlPort         string `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlDB           string `envconfig:"MYSQL_DB" default:"dev_database"`
	GrpcHost          string `envconfig:"GRPC_HOST" default:"artics-recommender-app"`
	GrpcPort          string `envconfig:"GRPC_PORT" default:"50051"`
	GCPServiceKeyJSON string `envconfig:"GCP_SERVICE_KEY_JSON" required:"true"`
	AWSProfile        string `envconfig:"AWSProfile" default:"di"`
	AWSS3Bucket       string `envconfig:"AWSS3Bucket" default:"bucket"`
	AWSS3Key          string `envconfig:"AWSS3Key" default:"key"`
}

// LoadEnvironment - loads the environment variables
func LoadEnvironment() (Environment, error) {
	env := Environment{}
	if err := envconfig.Process("", &env); err != nil {
		return env, xerrors.Errorf("Failed to LoadEnvironment: %w", err)
	}

	return env, nil
}
