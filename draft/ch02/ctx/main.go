package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	// 创建截止时间
	d := time.Now().Add(shortDuration)
	// 创建有截止时间的 Context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 主动让下游结束
	defer cancel()

	// 使用select 监听1秒和有截止时间的 Context 哪个先结束
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done(): // 被上游通知结束
		fmt.Println(ctx.Err())
	}
}
