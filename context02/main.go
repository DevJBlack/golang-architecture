package main

import (
	"context"
	"fmt"

	"github.com/DevJBlack/golang-architecture/tree/master/context02/session"
)

func main() {
	ctx := context.Background()
	ctx = session.SetUserId(ctx, 1)
	ctx = session.SetAdminAccess(ctx, true)

	uID := session.GetUserId(ctx)
	isAdmin := session.GetAdmin(ctx)
	fmt.Printf("Use %d is an admin %t", uID, isAdmin)
}
