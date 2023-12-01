package aferoKit

import (
	"github.com/spf13/afero"
	"regexp"
	"time"
)

var (
	NewOsFs func() afero.Fs = afero.NewOsFs

	NewMemMapFs func() afero.Fs = afero.NewMemMapFs

	NewHttpFs func(source afero.Fs) *afero.HttpFs = afero.NewHttpFs

	NewIOFS func(fs afero.Fs) afero.IOFS = afero.NewIOFS

	NewRegexpFs func(source afero.Fs, re *regexp.Regexp) afero.Fs = afero.NewRegexpFs

	NewReadOnlyFs func(source afero.Fs) afero.Fs = afero.NewReadOnlyFs

	NewCacheOnReadFs func(base afero.Fs, layer afero.Fs, cacheTime time.Duration) afero.Fs = afero.NewCacheOnReadFs

	NewCopyOnWriteFs func(base afero.Fs, layer afero.Fs) afero.Fs = afero.NewCopyOnWriteFs

	NewBasePathFs func(source afero.Fs, path string) afero.Fs = afero.NewBasePathFs
)
