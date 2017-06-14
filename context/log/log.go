package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const reqestIDKey = key(42)

// Println ...
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(reqestIDKey).(int64)
	if !ok {
		log.Println("could not find ID in context")
		return
	}
	log.Printf("%d: %s", id, msg)
}

// Decorate ...
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, reqestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
