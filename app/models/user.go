/*
 * @Author: Jadedever
 * @Date: 2022-06-24 16:30:07
 * @LastEditors: Jadedever
 * @LastEditTime: 2022-06-24 17:09:53
 * @FilePath: /goravel/app/models/user.go
 * @Description:
 *
 * Copyright (c) 2022 by Jadedever, All Rights Reserved.
 */
package models

import "github.com/goravel/framework/database/orm"

type User struct {
	orm.Model
	orm.SoftDeletes

	Name     string
	Username string
	Password string
}
