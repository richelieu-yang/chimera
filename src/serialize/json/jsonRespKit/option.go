package jsonRespKit

type (
	options struct {
		filePaths []string

		fileDataSlice []*FileData
	}

	Option func(opts *options)
)

func loadOptions(optionSlice ...Option) *options {
	opts := &options{
		filePaths:     nil,
		fileDataSlice: nil,
	}

	for _, option := range optionSlice {
		option(opts)
	}

	return opts
}

// WithFilePaths 存储code和msg对应关系的文件（路径）
func WithFilePaths(filePathSlice []string) Option {
	return func(opts *options) {
		opts.filePaths = filePathSlice
	}
}

// WithFileDataSlice 存储code和msg对应关系的文件（类型、内容）
func WithFileDataSlice(fileDataSlice []*FileData) Option {
	return func(opts *options) {
		opts.fileDataSlice = fileDataSlice
	}
}
