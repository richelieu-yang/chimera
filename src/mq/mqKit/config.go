package mqKit

import "github.com/apache/rocketmq-clients/golang/v5/credentials"

type (
	Config struct {
		Endpoints   []string                        `json:"endpoints"`
		Credentials *credentials.SessionCredentials `json:"credentials"`
	}
)
