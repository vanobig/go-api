package main

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/users",
		GetUsers,
	},
	Route{
		"GET",
		"/users/:id",
		GetUser,
	},
	Route{
		"POST",
		"/users",
		CreateUser,
	},
	Route{
		"PUT",
		"/users/:id",
		UpdateUser,
	},
	Route{
		"DELETE",
		"/users/:id",
		DeleteUser,
	},
}
