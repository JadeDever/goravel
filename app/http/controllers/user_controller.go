/*
 * @Author: Jadedever
 * @Date: 2022-06-24 16:30:07
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-07-05 22:45:42
 * @FilePath: /goravel/app/http/controllers/user_controller.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package controllers

import (
	"goravel/app/models"

	"github.com/gin-gonic/gin"
	"github.com/goravel/framework/support/facades"
)

type UserController struct {
}

func (r UserController) Show(ctx *gin.Context) {

	var users []models.User
	facades.DB.Find(&users, []int{1, 2, 30})
	facades.Response.Success(ctx, users)
}

func (controller UserController) Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{}
	facades.DB.Where(&models.User{Username: username, Password: password}).First(&user)
	facades.Response.Success(c, user)
}

func (controller UserController) Register(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	user := models.User{Name: name, Username: username, Password: password}
	facades.DB.Create(&user)

	facades.Response.Success(c, user)
}
