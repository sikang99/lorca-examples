package main

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"time"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 720, 320)
	if execErr, ok := err.(*exec.Error); ok {
		lorca.PromptDownload()
		log.Fatalf("\nChrome could not be started. Do you have Chrome installed? %s\n", execErr)
	} else if err != nil {
		log.Fatalf("\nOps. Something went wrong while starting the application: %s", err)
	}
	defer ui.Close()

	ui.Load("data:text/html," + url.PathEscape(`
	<html>
	  <head><title>Calling Go from Js and Js from Go</title></head>
	  <body>
	  <h1>1 - Calling Go code from page's Javascript</h1>
	  <script>
	  function incrementFromGoCode(button) {
        // add() exists in Js because we called ui.Bind("add", func(){}) in Go
		add(parseInt(button.value), 2).then(function(valueFromGo){
			button.value = valueFromGo;
		});
	  }
	  </script>
	  Click to call Go: <input type="button" id="button1" value="0" onclick="incrementFromGoCode(this)">
	  <h1>2 - Calling Javascript from Go</h1>
	  <p>Value updated from Go: <span id="span1">0</span></p>
	  </body>
	</html>`))

	// Allows Js to call our Go function add()
	ui.Bind("add", func(a, b int) int { return a + b })

	// Calls Js to increment value in HTML every second
	go func() {
		for i := 0; ; i++ {
			ui.Eval(fmt.Sprintf("document.getElementById('span1').innerHTML = %d;", i))
			time.Sleep(1 * time.Second)
		}
	}()

	<-ui.Done()
}
