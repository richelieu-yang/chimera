package logrusKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath("ccc.log"))
	assert.Nil(t, err)

	NewLogger()

	type args struct {
		options []LoggerOption
	}
	tests := []struct {
		name string
		args args
		want *logrus.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
