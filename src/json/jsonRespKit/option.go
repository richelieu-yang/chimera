package jsonRespKit

type (
	options struct {
		// filePathSlice
		filePathSlice []string

		// fileDataSlice
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

func WithFilePathSlice(filePathSlice []string) Option {
	return func(opts *options) {
		opts.filePathSlice = filePathSlice
	}
}

func WithFileDataSlice(fileDataSlice []*FileData) Option {
	return func(opts *options) {
		opts.fileDataSlice = fileDataSlice
	}
}
