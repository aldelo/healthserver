package main

import (
	ginw "github.com/aldelo/common/wrapper/gin"
	"github.com/aldelo/common/wrapper/gin/ginbindtype"
	"github.com/aldelo/common/wrapper/gin/ginhttpmethod"
	"github.com/aldelo/common/wrapper/systemd"
	ws "github.com/aldelo/connector/webserver"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

/*
 * Copyright 2020-2021 Aldelo, LP
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
	Windows Server Service Install => c:\sc.exe create xyzService binpath= c:\xyzFolder\xyz.exe type= own start= auto
	Linux Ubuntu Service Install => sudo systemctl start xyzService.service (See comments in common package /systemd wrapper)
 */
func main() {
	svc := &systemd.ServiceProgram{
		ServiceName: "HealthServer",
		DisplayName: "Health Server",
		Description: "Provides Http Endpoint for Health Check Service",
		StartServiceHandler: startServiceHandler,
		StopServiceHandler: nil,
	}

	svc.Launch()
}

func startServiceHandler(port int) {
	g := ws.NewWebServer("HealthServer", "server", "")

	g.Routes = map[string]*ginw.RouteDefinition{
		"base": {
			Routes: []*ginw.Route{
				{
					RelativePath: "/test",
					Method: ginhttpmethod.GET,
					Binding: ginbindtype.UNKNOWN,
					BindingInputPtr: nil,
					Handler: func(c *gin.Context, bindingInputPtr interface{}) {
						c.Status(200)
					},
				},
			},
			CorsMiddleware: &cors.Config{},
		},
	}

	if err := g.Serve(); err != nil {
		log.Println("Error: " + err.Error())
	} else {
		log.Println("Run OK")
	}
}
