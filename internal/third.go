package internal

import (
	jose "github.com/dvsekhvalnov/jose2go"
	"golang.org/x/arch/x86/x86asm"
	"golang.org/x/crypto/cast5"
	"golang.org/x/exp/slog"
	"golang.org/x/image/tiff"
	"golang.org/x/mobile/bind"
	"golang.org/x/mod/modfile"
	"golang.org/x/net/proxy"
	"golang.org/x/oauth2/jws"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/execabs"
	"golang.org/x/term"
	"golang.org/x/text/currency"
	"golang.org/x/time/rate"
	"golang.org/x/tools/blog"
)

func init() {
	/*
		TODO: 看后续 github.com/apache/pulsar-client-go v0.11.1 会不会更新.
		Richelieu: 为了引用最新的 dvsekhvalnov/jose2go 库（v1.6.0），v1.5.0有脆弱性漏洞，
	*/
	_ = jose.DEF

	{
		var _ x86asm.Arg
		var _ *cast5.Cipher
		var _ *bind.Generator
		var _ *modfile.Comment
		var _ jws.Signer
		var _ *errgroup.Group
		var _ *execabs.Error
		var _ *term.Terminal
		var _ *blog.Doc
		var _ slog.Handler
		var _ *tiff.Options
		var _ proxy.Dialer
		var _ *currency.Amount
		var _ *rate.Limit
	}
}
