package aferoKit

import (
	"github.com/spf13/afero"
	"regexp"
	"time"
)

var (
	// NewOsFs 创建一个新的 "基于操作系统" 的文件系统实例.
	NewOsFs func() afero.Fs = afero.NewOsFs

	// NewMemMapFs 创建一个新的 "基于内存" 的文件系统实例.
	NewMemMapFs func() afero.Fs = afero.NewMemMapFs

	NewHttpFs func(source afero.Fs) *afero.HttpFs = afero.NewHttpFs

	// NewIOFS 返回一个实现了afero.Fs接口的IOFS对象（这个IOFS对象是对Go1.16的io/fs.FS文件系统抽象的包装）.
	NewIOFS func(fs afero.Fs) afero.IOFS = afero.NewIOFS

	NewRegexpFs func(source afero.Fs, re *regexp.Regexp) afero.Fs = afero.NewRegexpFs

	NewReadOnlyFs func(source afero.Fs) afero.Fs = afero.NewReadOnlyFs

	NewCacheOnReadFs func(base afero.Fs, layer afero.Fs, cacheTime time.Duration) afero.Fs = afero.NewCacheOnReadFs

	NewCopyOnWriteFs func(base afero.Fs, layer afero.Fs) afero.Fs = afero.NewCopyOnWriteFs

	NewBasePathFs func(source afero.Fs, path string) afero.Fs = afero.NewBasePathFs
)
