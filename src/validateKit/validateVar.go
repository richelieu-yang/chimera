package validateKit

import "context"

// Var 验证变量.
func Var(field interface{}, tag string) error {
	v := New()
	return v.Var(field, tag)
}

func VarCtx(ctx context.Context, field interface{}, tag string) error {
	v := New()
	return v.VarCtx(ctx, field, tag)
}

// VarWithValue 验证变量，同时验证另一个变量.
func VarWithValue(field interface{}, other interface{}, tag string) error {
	v := New()
	return v.VarWithValue(field, other, tag)
}

func VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) error {
	v := New()
	return v.VarWithValueCtx(ctx, field, other, tag)
}
