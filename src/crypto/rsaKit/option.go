package rsaKit

type (
	rsaOptions struct {
		// format 密钥格式（私钥）
		format KeyFormat
		// password 私钥密码（可以为""）
		password string
	}

	RsaOption func(opts *rsaOptions)
)

func loadOptions(options ...RsaOption) *rsaOptions {
	opts := &rsaOptions{
		format:   PKCS8,
		password: "",
	}
	for _, option := range options {
		option(opts)
	}

	// check
	switch opts.format {
	case PKCS1:
		fallthrough
	case PKCS8:
		// do nothing
	default:
		opts.format = PKCS8
	}

	return opts
}

// WithFormat 配置: 密钥格式（私钥）
func WithFormat(format KeyFormat) RsaOption {
	return func(opts *rsaOptions) {
		opts.format = format
	}
}

// WithPassword 配置: 私钥密码（可以为""）
func WithPassword(password string) RsaOption {
	return func(opts *rsaOptions) {
		opts.password = password
	}
}
