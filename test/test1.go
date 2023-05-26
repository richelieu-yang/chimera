package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	data, err := fileKit.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}

	data, err = rsaKit.Decrypt([]byte("D/VGDoQKkFOaXBLBy368+1LvDHcdgaANCiMFGn7+wy09Rzkz0HS3gz+V5KNLRrI3ChULRrkZVhqQ/0Ngq8nUX1on+a/m5A2uJCVRLKUPilZsFnjstjMOK3v31lMC0tGB86P+zuAR0kSPs0YDkXEy3pLlGulb1Ezh77zQ4ACVMI3Ywuh1/hPrJWB9WT2AlW5IoXkQrdtf6SJvG7xGf1uQga9H5nPOxbozsgKzRQRmTDyFUVTnIIKoDDX3K7i5ADIYoAmZX2fGk0FDXFUImaNRexBfpduCEyYI5UG6RzQ6a00K2Z3Oxf3xKYZWkWKrFYp884sG8V5D6EBxfUmMUOfOaQ=="),
		data, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
