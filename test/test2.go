package main

import (
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/yamlKit"
)

func main() {
	str := `{"ip":"127.0.0.1","debug":true,"LogValidityPeriod":3,"gin":{"mode":"debug","hostName":"","port":9942,"colorful":true,"pprof":false,"ssl":{"access":false,"certFile":"","keyFile":""},"middleware":{"bodyLimit":-1,"gzip":false,"xFrameOptions":"","cors":{"access":false,"origins":null},"referer":null}},"redis":{"userName":"","password":"D/VGDoQKkFOaXBLBy368+1LvDHcdgaANCiMFGn7+wy09Rzkz0HS3gz+V5KNLRrI3ChULRrkZVhqQ/0Ngq8nUX1on+a/m5A2uJCVRLKUPilZsFnjstjMOK3v31lMC0tGB86P+zuAR0kSPs0YDkXEy3pLlGulb1Ezh77zQ4ACVMI3Ywuh1/hPrJWB9WT2AlW5IoXkQrdtf6SJvG7xGf1uQga9H5nPOxbozsgKzRQRmTDyFUVTnIIKoDDX3K7i5ADIYoAmZX2fGk0FDXFUImaNRexBfpduCEyYI5UG6RzQ6a00K2Z3Oxf3xKYZWkWKrFYp884sG8V5D6EBxfUmMUOfOaQ==","minIdleConns":64,"maxIdleConns":256,"poolSize":512,"mode":0,"singleNodeConfig":{"addr":"127.0.0.1:6379","db":0},"masterSlaverConfig":{},"sentinelConfig":{"masterName":"mymaster","sentinelAddrs":["172.18.21.17:26380","172.18.21.17:26381","172.18.21.17:26382"],"db":0},"clusterConfig":{"addrs":["192.168.1.25:6380","192.168.1.25:6381","192.168.1.25:6382","192.168.1.25:6383","192.168.1.25:6384","192.168.1.25:6385"]}},"pulsar":{"addresses":["192.168.80.27:6650","192.168.80.42:6650","192.168.80.43:6650"],"verify":{"topic":"test","print":true}}}`
	m := map[string]interface{}{}
	if err := jsonKit.UnmarshalFromString(str, &m); err != nil {
		panic(err)
	}

	if err := yamlKit.MarshalToFileWithJsonTag(m, "aaa.yaml"); err != nil {
		panic(err)
	}

	//v := viper.New()
	//v.SetConfigType("json")
	//if err := v.ReadConfig(ioKit.NewReader([]byte(str))); err != nil {
	//	panic(err)
	//}
	//
	//if err := v.WriteConfigAs("a.yaml"); err != nil {
	//	panic(err)
	//}
}
