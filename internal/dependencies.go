package internal

import (
	jose "github.com/dvsekhvalnov/jose2go"
	"golang.org/x/arch/x86/x86asm"
	"golang.org/x/crypto/cast5"
	"golang.org/x/exp/slog"
	"golang.org/x/mobile/bind"
	"golang.org/x/mod/modfile"
	"golang.org/x/oauth2/jws"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/execabs"
	"golang.org/x/term"
	"golang.org/x/tools/blog"
)

func init() {
	/*
		TODO: 看后续 github.com/apache/pulsar-client-go v0.11.1 会不会更新.
		Richelieu: 为了引用最新的 dvsekhvalnov/jose2go 库（v1.6.0），v1.5.0有脆弱性漏洞，
	*/
	_ = jose.DEF

	{
		// golang.org/x/arch
		var _ x86asm.Arg
		// golang.org/x/crypto
		var _ *cast5.Cipher
		// golang.org/x/mobile
		var _ *bind.Generator
		// golang.org/x/mod
		var _ *modfile.Comment
		// golang.org/x/oauth2
		var _ jws.Signer
		// golang.org/x/sync
		var _ *errgroup.Group
		// golang.org/x/sys
		var _ *execabs.Error
		// golang.org/x/term
		var _ *term.Terminal
		// golang.org/x/tools
		var _ *blog.Doc
		// golang.org/x/exp
		var _ slog.Handler
	}
}
