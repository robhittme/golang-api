package main

import (
	w "github.com/robhittme/golang-api/webserver"
	m "github.com/robhittme/golang-api/model"
)

func main() {
	m.Init()
	w.InitRoutes()

}
