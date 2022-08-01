package main

import (
	"context"
	"kafka/lib"
)

var topics1 = []string{"datta", "test"}

func main1() {
	// create a new context
	var s string = "Hello World"
	sb := []byte(s)
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	lib.CreateTopics(ctx, topics1)
	go lib.KF_Produce(ctx, sb, "test")
	lib.KF_Consume(ctx, "datta")
	// lib.KF_Consume(ctx, "test")

}
