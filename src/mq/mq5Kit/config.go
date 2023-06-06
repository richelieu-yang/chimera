package mq5Kit

import "github.com/apache/rocketmq-clients/golang/v5/credentials"

type (
	Config struct {
		Endpoint    string
		Credentials *credentials.SessionCredentials `json:"credentials"`
	}
)
