package main

import (
	"context"
	"fmt"
)

// 包内私有类型，防止其他地方修改 context 中的key
type authkay struct{}

// 存储的 auth 信息，可定义在 公共包里
type AuthUser struct {
	Id   string
	Name string
}

// 写入ctx
func NewContext(ctx context.Context, u AuthUser) context.Context {
	return context.WithValue(ctx, authkay{}, u)
}

// 读取 ctx
func FromContext(ctx context.Context) (tr AuthUser, ok bool) {
	tr = ctx.Value(authkay{}).(AuthUser)
	return tr, true
}

// example:
func main() {
	ctx := context.Background()

	// middleware  中间件中解析 token 到 用户 给 ctx 写入
	ctx = NewContext(ctx, AuthUser{Id: "1", Name: "kingofzihua"})

	// 在controller 中 从 ctx 中读取 auth 信息
	auth, ok := FromContext(ctx)
	if !ok {
		panic("auth error")
	}

	// 使用 auth user
	fmt.Println(auth)

}
