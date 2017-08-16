package main

import (
	"github.com/codegangsta/martini"
)
func GetApi() *martini.ClassicMartini {
	api := martini.Classic()
	return api
}
