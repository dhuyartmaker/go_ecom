package router

import "github.com/go-ecommerce/internal/router/user"

type Router struct {
	UserRouter user.UserRouterGroup
}

var RouterApp = new(Router)