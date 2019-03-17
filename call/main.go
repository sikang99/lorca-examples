package main

import (
	"net/url"

	"github.com/zserge/lorca"
)

var ui lorca.UI

func main() {
	ui, _ = lorca.New("", "", 480, 320)
	defer ui.Close()

	// GoのfuncをJavaScriptで呼び出せるようBind
	ui.Bind("helloFromGo", helloFromGo)

	// 今回はurl.PathEscapeでそのままHTMLを配置
	// ※下記はHTTPサーバをたててそこから取得する例
	//   https://github.com/zserge/lorca/tree/master/examples/counter
	ui.Load("data:text/html," + url.PathEscape(`
    <!doctype html>
    <html lang="ja">
    <head>
    <title>Lorca Sample</title>
	<script>
        function helloFromJavaScript() {
            return 'Hello from JavaScript!'
        }
    </script>
    </head>
    <body>
        <button onclick="helloFromGo()"> Press Me! </button>
        <div id="content"></div>
    </body>
    </html>
    `))

	<-ui.Done()
}

func helloFromGo() {
	// JavaScriptのfunction呼び出し
	msg := ui.Eval(`helloFromJavaScript();`).String() + "<br>" + "Hello From Go!"
	// HTMLに反映
	ui.Eval(`document.getElementById('content').innerHTML="` + msg + `";`)
}
