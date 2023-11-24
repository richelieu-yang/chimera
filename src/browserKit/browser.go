package browserKit

import (
	"github.com/pkg/browser"
	"io"
)

var (
	// OpenURL Open a new browser window pointing to url.
	OpenURL func(url string) error = browser.OpenURL

	// OpenFile Open new browser window for the file path.
	OpenFile func(path string) error = browser.OpenFile

	// OpenReader Consume the contents of r and presents the results in a new browser window.
	OpenReader func(r io.Reader) error = browser.OpenReader
)

// SetOutputs
/*
@param stdout 可以为nil（默认: os.Stdout）
@param stderr 可以为nil（默认: os.Stderr）
*/
func SetOutputs(stdout, stderr io.Writer) {
	if stdout != nil {
		browser.Stdout = stdout
	}
	if stderr != nil {
		browser.Stderr = stderr
	}
}
