package jsonRespKit

type (
	options struct {
		filePathSlice []string

		fileDataSlice []*FileData
	}

	Option func(opts *options)
)

func loadOptions(optionSlice ...Option) *options {
	opts := &options{
		filePathSlice: nil,
		fileDataSlice: nil,
	}

	for _, option := range optionSlice {
		option(opts)
	}

	return opts
}

// WithFilePathSlice 存储code和msg对应关系的文件（路径）
func WithFilePathSlice(filePathSlice []string) Option {
	return func(opts *options) {
		opts.filePathSlice = filePathSlice
	}
}

// WithFileDataSlice 存储code和msg对应关系的文件（类型、内容）
func WithFileDataSlice(fileDataSlice []*FileData) Option {
	return func(opts *options) {
		opts.fileDataSlice = fileDataSlice
	}
}
