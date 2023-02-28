package urlKit

import "net/url"

func Parse(rawURL string) (*url.URL, error) {
	return url.Parse(rawURL)
}
