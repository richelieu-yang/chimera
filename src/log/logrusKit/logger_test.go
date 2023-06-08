package logrusKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	/* 输出到控制台 */
	logger := NewLogger()
	logger.Info("to console")

	/* 输出到文件（可rotate） */
	path := "test.log"
	writer, err := ioKit.NewLumberjackWriteCloser(path)
	assert.Nil(t, err)
	logger = NewLogger(WithWriter(writer))
	logger.Info("to rotatable file")

	err = fileKit.AssertExistAndIsFile(path)
	assert.Nil(t, err)
}

func TestNewFileLogger(t *testing.T) {
	type args struct {
		path    string
		options []LoggerOption
	}
	tests := []struct {
		name    string
		args    args
		want    *logrus.Logger
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFileLogger(tt.args.path, tt.args.options...)
			if !tt.wantErr(t, err, fmt.Sprintf("NewFileLogger(%v, %v)", tt.args.path, tt.args.options...)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewFileLogger(%v, %v)", tt.args.path, tt.args.options...)
		})
	}
}
