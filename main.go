package main

import (
	docs "github.com/agustinlare/tf-varset-updater/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	TerraformApi = "https://app.terraform.io/api/v2"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		tfc := v1.Group("/")
		{
			tfc.GET("/listvariablesets", ListVariableSet)
			tfc.GET("/showvariableset/:varset_id", ShowVariableSet)
			tfc.GET("/showvariable/:var_id", ShowVariable)
			tfc.POST("/appendvariable/:varset_id", AppendVariable)
			tfc.PATCH("/updatevariable/:varset_id", UpdateVariable)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
