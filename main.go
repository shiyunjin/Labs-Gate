package main

import (
	"gitee.com/shiyunjin/SchoolNetwork/system"
)

func main() {
	r := router.Router()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
