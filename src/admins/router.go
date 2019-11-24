package admins

import "github.com/gin-gonic/gin"

////////// Public Methods //////////

// Register will register routers
func Register(r *gin.RouterGroup) {
	r.GET("/admin", createHandler)
}

////////// Private Methods //////////
