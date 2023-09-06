package rocketmq5Kit

import "github.com/apache/rocketmq-clients/golang/v5/credentials"

type (
	Config struct {
		Endpoints      []string                        `json:"endpoints" yaml:"endpoints"`
		Credentials    *credentials.SessionCredentials `json:"credentials,optional" yaml:"credentials"`
		ValidatedTopic string                          `json:"validatedTopic,optional" yaml:"validatedTopic"`
		ClientLogPath  string                          `json:"clientLogPath,optional" yaml:"clientLogPath"`
	}
)
