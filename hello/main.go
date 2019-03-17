package main

import (
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("", "", 480, 320)
	defer ui.Close()

	ui.Load("data:text/html," + url.PathEscape(`
	<!DOCTYPE html>
	<html>
	<head></head>
	<body><h1>Hello Lorca with go</h1></body>
	</html>`))

	<-ui.Done()
}
