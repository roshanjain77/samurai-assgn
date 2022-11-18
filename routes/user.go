package routes

import (
	"login-signup-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/getusers", controllers.GetUsers)
	router.POST("/register", controllers.Register)
	router.DELETE("/deleteuser/:id", controllers.DeleteUser)
	router.PUT("/updateuser/:id", controllers.UpdateUser)
	router.POST("/login", controllers.Login)
	router.GET("/find/:email", controllers.FindUserByEmail)
}

func DashboardRoute(router *gin.Engine) {
	router.POST("/:id/createdashboard", controllers.CreateDashboard)
	router.GET("/:id/getdashboard", controllers.GetDashboards)
	router.GET("/addwidget/:id/:widget", controllers.AddWidget)
}

func ManageAccess(router *gin.Engine){
	router.POST("/changeRole", controllers.ChangeRole)
	router.GET("/getPermissions/:id", controllers.GetPermissions)
}
