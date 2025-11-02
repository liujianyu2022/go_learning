package main

import "context"

// 定义所有 context key
type contextKey struct {
	name string
}

var (
	key2 = contextKey{"key2"}
)

func main() {
	ctx := context.WithValue(context.Background(), "key1", "value1")			// 建议不要使用内置类型（如 string）作为 context 的 key，以避免潜在的键冲突。
	ctx = context.WithValue(ctx, key2, "value2")

	if value, ok := ctx.Value("key1").(string); ok {
		println("key1 =", value)
	}
	if value, ok := ctx.Value(key2).(string); ok {
		println("key2 =", value)
	}
}
