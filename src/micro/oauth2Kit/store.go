package oauth2Kit

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/store"
)

// NewMemoryTokenStore
/*
PS: 使用BuntDB，存储到内存中.
*/
var NewMemoryTokenStore func() (oauth2.TokenStore, error) = store.NewMemoryTokenStore
