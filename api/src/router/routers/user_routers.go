package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Router{
	createUser,
	getUser,
	listUsers,
	updateUser,
	deleteUser,
}

var (
	createUser = Router{
		Uri:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	}
	getUser = Router{
		Uri:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequired: false,
	}
	listUsers = Router{
		Uri:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.ListUsers,
		AuthRequired: false,
	}
	updateUser = Router{
		Uri:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	}
	deleteUser = Router{
		Uri:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	}
)
