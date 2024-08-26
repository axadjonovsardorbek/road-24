package main

import (
	"mobile/app"
	cf "mobile/config"
)

func main(){
	config := cf.Load()

	app.Run(config)
}