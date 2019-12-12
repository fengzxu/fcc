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

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"middleware"
	"net/http"
	"routers"
	"service"
	"time"
)

func main() {
	//initlog()

	log.SetPrefix("[fcc-server] ")
	//init system
	err := service.InitOnSystemStart()
	if err != nil {
		log.Println("error on init system:", err.Error())
		return
	}
	//init http
	inithttp()
	//println("service started.")
}

func inithttp() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	// recover
	app.Use(middleware.RecoveryMiddleware())

	// routers
	routers.RegisterRouter(app)

	//go initHTTPServer(app)
	initHTTPServer(app)

}

func initHTTPServer(handler http.Handler) {
	port := ":1206"
	srv := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("Server started on ", srv.Addr)
	srv.ListenAndServe()
}
