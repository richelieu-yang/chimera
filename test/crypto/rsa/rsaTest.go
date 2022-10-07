package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/crypto/rsaKit"
)

func main() {
	str := "异世界新人入坑指南\n\nedit：1区催眠大师，感谢1群客服A和老板娘的协助\n比较实用的礼包码：\nrtrusdmqyt  王国奴隶卡池 10连一发，500钻石，玩具熊x5\nertjmxuqyy 超级电池x10 (合计200堕落值)\nqwegfwqybc 500钻石，蚊香x3，技能法典x3，钻戒x5(100爱情值)\nrtyfmkadny 鑽石*1000（新增）\n\n端午节、六一、520这种节日礼包码一般都只有1-2天，一般群里和游戏公告都有，过期就领不了了，所以节假日的时候自己关注一下公告或者群里。\n\n4-40怎么过？\n早期英雄比较少，4-40之前是一个节点，这个时候一般玩家手里有的 丽萨（猫猫），杰西嘉（女仆），夏洛特（大公主），菲欧娜（大奶法师，3-40主线给碎片记得合成），奥莉娅（黑皮精灵3-40进地牢），多莉丝（小鹿，首充英雄，如"
	//pubKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqklW9t+7WXEgJ2kYgl6A\nHywTy0qLq36DsOia9IEKt3ko0uZeLjXFyi0OpzOdoMEdQW/Kwbx416tpmLPKLrEI\nJjd0GdbA6XrwSTU4dTDoU//8Ret6TxzjhG83YguyY6ORQ6zP0UkxvprKXOBtWIHy\nHXTinXDjXt7oYW02bcsckcZQ9OD72nF6R3yE2xMoz/wpNMxPVrOiMjZPodVuQZVH\nD/LNv6xzrunPr8EshDztYAhWZornRPsLPLgvgotkcJXo5ExCYAGn1zUcJs3DKsZi\nf50oBX/WPTz4aKzmtKalMpycDGyP4PRijzfjAz7IJFksjs55ExEa1xM77/A3JKTe\nSQIDAQAB\n-----END PUBLIC KEY-----\n"
	//priKey := "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCqSVb237tZcSAn\naRiCXoAfLBPLSourfoOw6Jr0gQq3eSjS5l4uNcXKLQ6nM52gwR1Bb8rBvHjXq2mY\ns8ousQgmN3QZ1sDpevBJNTh1MOhT//xF63pPHOOEbzdiC7Jjo5FDrM/RSTG+mspc\n4G1YgfIddOKdcONe3uhhbTZtyxyRxlD04PvacXpHfITbEyjP/Ck0zE9Ws6IyNk+h\n1W5BlUcP8s2/rHOu6c+vwSyEPO1gCFZmiudE+ws8uC+Ci2RwlejkTEJgAafXNRwm\nzcMqxmJ/nSgFf9Y9PPhorOa0pqUynJwMbI/g9GKPN+MDPsgkWSyOznkTERrXEzvv\n8DckpN5JAgMBAAECggEAfjaztXGozCdz/KIB34cpkXE7Dp+nHCo+c55EU9tdVRYB\ncyf6QGfsvOhehwFT4PWhuIGil6wZKUlMkDkaLzdZ4fGSJSCuhZ9wf+JpgTsJyFDg\n6/YYBpFT9TH1q78dropqve1VbzrAGxVhzisuAL739GtRF/63F/dB6AxpnkFAYK6W\nLP5v7/BDBfU7Mp/7ivc6lS/Aqd3W2QJ5OkUdhcp4ztWP3UstOFmrNpvCnZsbt6dA\nABE58+Hs5o1LXk4UUX9N7Fex4uzoicYqASiysd4ewR6/2AAVabnMa7aC9odRTjM4\nlF8fE5B5qaAFxclYVXFSOiwMlKXESLybesnDPwoAPQKBgQDbLNn9P3P7qVvtC5z9\nqWb1iTDqT2EVtFgglz3mmGw/M0aeKgmqiSoF6f2i73828WjLd19n+hpNghzHF4p1\nIrKZcn3YaqZheXyjfvMpSNDNH0lfkhUZSi6YmxdZFtaXUcC6kIKWfhHwh7cZ76Bq\nqgS5DMhJWsF/9MohmYwqkaj4OwKBgQDG5bCnOTi/ZzHWFtBbS1VILgWJiuVIgcLJ\nunedW9NMFUh8XeiXlBTXSgOuCyn3U0rIraI+y5JSWfH+D+qaBiryKfL6M1coCarw\nSEJh9nzbEXIOI9kL5eaOnZVkOrc0wljRPPff5V3V32tppm4S27LbqQhhTxLh0s1i\nP00Ame8fSwKBgQDY3h1ob3PFDUQPXpFdkVR2GeSTxC2tO0CEJwKx/BoHwyXZ9ICS\nBJzXMzusEEyRQffOc/SsPpXuIZN2ED4JP2b3XIlXLdgp1PrWEBRswkcQ4CdK8JH9\n7yACElvc5DM7kUIhmEIOxPndRXyQpkHVqRt5O0OnCa0zfjTPYE6IBI7r5wKBgQCu\nm+pvg76b2b5YkoKsAwCFc1GNJBCRjMkyS67CyGBBBNYzyEVu/KMnNWTxn/9hjw+y\nCSzML+7uuOPHBLYSLaP6IHWczQGQtpwIgYc1dSy7LR+R5kqfNV8oiJMkAJ/hg1Rx\nrDJ3rmaoYZFsmJ09lP359PTUaJTNB0EjJSs2gez47QKBgF5xIJUD3qGe+2CFlrfU\nRvoe4yC0ukuhiJfJNE+5i72xpg4i8eQA+Cl9FwnnSLtT9VhcWS995a1Fxs2/vFz7\n7q5E7HabLYKJ3l8AUrTLJrOJLRmytciOFnSvH+GpANkE0JSlRioGCf45GE1d6SS+\nJEQzB6mxPAX0hMI0qVdkQhaY\n-----END PRIVATE KEY-----\n"

	str, err := rsaKit.EncryptToString(rsaKit.DefaultPublicKey, []byte(str))
	if err != nil {
		panic(err)
	}
	fmt.Println("密文：\n" + str)

	str, err = rsaKit.DecryptToString(rsaKit.DefaultPrivateKey, rsaKit.DefaultPassword, []byte(str))
	if err != nil {
		panic(err)
	}
	fmt.Println("原文：\n" + str)
}
