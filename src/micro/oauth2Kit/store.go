package oauth2Kit

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/store"
)

var NewMemoryTokenStore func() (oauth2.TokenStore, error) = store.NewMemoryTokenStore
