package jsonRespKit

type (
	options struct {
		filePathSlice []string
		fileDataSlice []*FileData
	}

	Option func(opts *options)
)

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

func loadOptions(optionSlice ...Option) *options {
	opts := &options{}

	for _, option := range optionSlice {
		option(opts)
	}

	return opts
}
