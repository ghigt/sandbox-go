package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ghigt/sandbox-go/context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	fmt.Printf("value for foo is %v\n", ctx.Value("foo"))

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		log.Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
