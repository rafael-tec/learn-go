package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "a3bz49c0")

	bookHotel(ctx, "Hotel Fazenda")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")

	fmt.Printf("Name: %s\nToken: %s\n", name, token)
}
