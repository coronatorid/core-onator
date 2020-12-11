package testhelper

import "context"

// TestContext for testing purpose
type TestContext struct {
	ctx context.Context
}

// NewTestContext create new test context
func NewTestContext() *TestContext {
	return &TestContext{
		ctx: context.Background(),
	}
}

// Ctx ...
func (t *TestContext) Ctx() context.Context {
	return t.ctx
}

// Get ...
func (t *TestContext) Get(key string) interface{} {
	return t.ctx.Value(key)
}

// Set ...
func (t *TestContext) Set(key string, val interface{}) {
	t.ctx = context.WithValue(t.ctx, key, val)
}
