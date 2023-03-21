package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/mapKit"
)

func SetTag(pMsg *pulsar.ProducerMessage, tag string) {
	m := pMsg.Properties

	mapKit.Set()
}

func GetTag(msg pulsar.Message) string {

}
