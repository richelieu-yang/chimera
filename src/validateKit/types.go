package validateKit

type (
	// Validatable 参考: "github.com/ardielle/ardielle-go/rdl"中的 rdl.Validatable
	Validatable interface {
		Validate() error
	}
)
