package baiduOcrKit

import (
	"github.com/sirupsen/logrus"
	"log"
	"testing"
	"time"
)

func TestRecognizeUniversalWords(t *testing.T) {
	err := SetApiKeyAndSecretKey("9Rryr5apq6FKGp3kAMR90d7D", "3LkHdrNyg10GlGSzKvyKCbkqc0gLV16b")
	if err != nil {
		log.Panic(err)
	}

	imgs := []string{
		"/Users/richelieu/Downloads/1.png",
		"/Users/richelieu/Downloads/2.png",
		//"/Users/richelieu/Downloads/3.png",
	}
	for _, img := range imgs {
		go func(img string) {
			words, err := RecognizeUniversalWords(img)
			if err != nil {
				logrus.Error(err)
				return
			}
			for _, wordResult := range words.WordsResults {
				logrus.Info(wordResult.Words)
			}
		}(img)
	}

	time.Sleep(time.Second * 3)
}
