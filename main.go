package main

import (
	"context"
	"kafka/lib"
)

var topics = []string{"datta", "test", "mcst"}

func main() {
	// create a new context
	// var s string = "Hello World"
	// sb := []byte(s)
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	lib.CreateTopics(ctx, topics)
	// go lib.KF_Produce(ctx, sb, "test")
	// go lib.KF_Consume(ctx, "datta")
	lib.KF_Consume(ctx, "test")

}
