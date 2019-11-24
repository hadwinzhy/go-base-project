package routers

import "github.com/gin-gonic/gin"

////////// Public Method //////////

////////// Private Method //////////

func routerErrorHandler(r *gin.Engine) {
	router404Handler(r)
	router500Handler(r)
}

func router404Handler(r *gin.Engine) {
	r.NoRoute()
}

func router500Handler(r *gin.Engine) {

}
