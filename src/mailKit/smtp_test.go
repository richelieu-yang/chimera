package mailKit

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"testing"
)

func TestSendMail(t *testing.T) {
	sender := &SmtpConfig{
		Address:  "smtp.163.com:25",
		Host:     "smtp.163.com",
		Account:  "miro42@163.com",
		Password: "ZRZUBJXZOFIOBNXM",
		//NickName: "yjs",
	}
	err := InitializeSmtp(sender, 2)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1; i++ {
		time := timeKit.FormatCurrentTime()
		mail := NewMail("", []string{"yjs@yozosoft.com", "richelieu042@gmail.com"}, "主题"+time, []byte(intKit.FormatIntToString(i, 10)), nil, []string{"miro42@163.com"}, nil)
		err = SendMail(mail)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("-------------------------------------------------")
}
