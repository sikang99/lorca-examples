package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/zserge/lorca"
)

const html = `
<html>
  <head>
    <title>Check URL</title>
  </head>
  <body>
    url: <input id="url" value="https://fosdem.org"></input>
    <button id="check" onclick="checkURL(document.getElementById('url').value)">Check</button>
    <div id="status"></div>
  </body>
</html>
`

func main() {
	ui, err := lorca.New("data:text/html,"+url.PathEscape(html), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("checkURL", func(url string) {
		ui.Eval(`document.getElementById('status').innerText = 'checking...';`)
		log.Println(url)
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			//ui.Close()
			ui.Eval(`document.getElementById('status').innerText = 'error!';`)
			<-ui.Done()
		}
		//log.Println(res.StatusCode)
		if res.StatusCode == http.StatusOK {
			ui.Eval(`document.getElementById('status').innerText = 'Online';`)
		} else {
			ui.Eval(`document.getElementById('status').innerText = 'Offline';`)
		}
	})

	<-ui.Done()
}

/*
ui.Bind("checkURL", checkURL(url string))
func checkURL(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		ui.Eval(`document.QuerySelector('#status').innerText = 'Online';`)
	} else {
		ui.Eval(`document.QuerySelector('#status').innerText = 'Offline';`)
	}
}
*/
