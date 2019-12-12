/*
Copyright xujf000@gmail.com .2020. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routers

import (
	"ccservice"
	"github.com/gin-gonic/gin"
	"service"
	"strconv"
)

func RegisterRouter(app *gin.Engine) {
	//public frontend files
	//app.StaticFS("/html", http.Dir("public"))

	//registe all route path
	registerPath(app)
}

func registerPath(app *gin.Engine) {

	//root
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, "fcc server.")
	})
	app.GET("/api", func(c *gin.Context) {
		c.JSON(200, "fcc api server.")
	})

	//netcon
	preNetcon := app.Group("/api/netcon")
	preNetcon.POST("/create", func(c *gin.Context) {
		netconid := c.PostForm("netconid")
		applya := c.PostForm("applya")
		applyb := c.PostForm("applyb")
		addr := c.PostForm("addr")
		area, _ := strconv.Atoi(c.DefaultPostForm("area", "0"))
		balance, _ := strconv.Atoi(c.DefaultPostForm("balance", "0"))
		c.JSON(200, service.NetconCreate(netconid, applya, applyb, addr, area, balance))
	})
	preNetcon.POST("/tocc", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		netconid := c.PostForm("netconid")
		applya := c.PostForm("applya")
		applyb := c.PostForm("applyb")
		addr := c.PostForm("addr")
		area := c.PostForm("area")
		balance := c.PostForm("balance")
		c.JSON(200, service.NetconToCC(uuid, netconid, applya, applyb, addr, area, balance))
	})
	preNetcon.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, service.NetconGetAll())
	})

	//cc-netcon
	preCCNetcon := app.Group("/api/cc/netcon")
	preCCNetcon.GET("/querybynetconid", func(c *gin.Context) {
		netconid := c.Query("netconid")
		c.JSON(200, ccservice.NetconQueryByNetconid(netconid))
	})
	preCCNetcon.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, ccservice.NetconQueryAll())
	})
	preCCNetcon.GET("/queryallid", func(c *gin.Context) {
		c.JSON(200, ccservice.NetconGetAllCCid())
	})

	//estatebook
	preEstatebook := app.Group("/api/estatebook")
	preEstatebook.POST("/create", func(c *gin.Context) {
		bookid := c.PostForm("bookid")
		netconid := c.PostForm("netconid")
		taxid := c.PostForm("taxid")
		owner := c.PostForm("owner")
		addr := c.PostForm("addr")
		area, _ := strconv.Atoi(c.DefaultPostForm("area", "0"))
		c.JSON(200, service.EstateBookCreate(bookid, netconid, taxid, owner, addr, area))
	})
	preEstatebook.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, service.EstateBookGetAll())
	})
	preEstatebook.POST("/tocc", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		bookid := c.PostForm("bookid")
		owner := c.PostForm("owner")
		area := c.PostForm("area")
		addr := c.PostForm("addr")
		c.JSON(200, service.EstateBookToCC(uuid, bookid, owner, addr, area))
	})

	//cc-estatebook
	preCCEstatebook := app.Group("/api/cc/estatebook")
	preCCEstatebook.GET("/querybybookid", func(c *gin.Context) {
		bookid := c.Query("bookid")
		c.JSON(200, ccservice.EstateBookQueryByBookid(bookid))
	})
	preCCEstatebook.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, ccservice.EstateBookQueryAll())
	})

	//estatetax
	preEstatetax := app.Group("/api/estatetax")
	preEstatetax.POST("/create", func(c *gin.Context) {
		taxid := c.PostForm("taxid")
		taxer := c.PostForm("taxer")
		area, _ := strconv.Atoi(c.DefaultPostForm("area", "0"))
		tax, _ := strconv.Atoi(c.DefaultPostForm("tax", "0"))
		c.JSON(200, service.EstateTaxCreate(taxid, taxer, area, tax))
	})
	preEstatetax.POST("/tocc", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		taxid := c.PostForm("taxid")
		taxer := c.PostForm("taxer")
		area := c.PostForm("area")
		tax := c.PostForm("tax")
		c.JSON(200, service.EstateTaxToCC(uuid, taxid, taxer, area, tax))
	})
	preEstatetax.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, service.EstateTaxGetAll())
	})

	//cc-estatetax
	preCCEstatetax := app.Group("/api/cc/estatetax")
	preCCEstatetax.GET("/querybytaxid", func(c *gin.Context) {
		taxid := c.Query("taxid")
		c.JSON(200, ccservice.EstateTaxQueryByTaxid(taxid))
	})
	preCCEstatetax.GET("/queryall", func(c *gin.Context) {
		c.JSON(200, ccservice.EstateTaxQueryAll())
	})
	preCCEstatetax.GET("/queryallid", func(c *gin.Context) {
		c.JSON(200, ccservice.EstateTaxQueryAllid())
	})
}
