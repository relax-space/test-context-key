package main

import (
	"context"
	"fmt"
	"time"
)

type ContextType string

var ContextName ContextType = "ContextName"

func main() {
	test1()
	fmt.Println("test2 start...")
	time.Sleep(2 * time.Second)
	test2()
	fmt.Println("Done")
	time.Sleep(10000 * time.Second)
}

func test1() {
	ctx := set1(context.Background())
	fmt.Println(ctx.Value(ContextName))
	time.Sleep(2 * time.Second)
	go func(ctx1 context.Context) {
		fmt.Println(ctx1.Value(ContextName))
	}(ctx)
}

func test2() {
	ctx := set1(context.Background())
	ctx = set2(ctx)
	fmt.Println(ctx.Value(ContextName))
	time.Sleep(2 * time.Second)
	go func(ctx1 context.Context) {
		fmt.Println(ctx1.Value(ContextName))
	}(ctx)
}

func set1(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextName, "hello")
}

func set2(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextName, "world")
}
