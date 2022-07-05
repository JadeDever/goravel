/*
 * @Author: Jadedever
 * @Date: 2022-06-24 16:30:07
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-07-05 22:05:19
 * @FilePath: /goravel/routes/web.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package routes

import (
	"goravel/app/http/controllers"

	"github.com/gin-gonic/gin"
	"github.com/goravel/framework/support/facades"
)

func Web() {
	facades.Route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello": "Goravel",
		})
	})

	facades.Route.GET("/user", controllers.UserController{}.Show)

	facades.Route.POST("/auth/login", controllers.UserController{}.Login)
	facades.Route.POST("/auth/register", controllers.UserController{}.Register)

}
