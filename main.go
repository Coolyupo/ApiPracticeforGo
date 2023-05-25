package main

import (
	"gurubear/routers"
)

func main() {
	web := routers.InitRouter()
	web.Run(":8080")
}
