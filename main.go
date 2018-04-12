package main

import (
	"github.com/kataras/iris"
	"net/http"
)

// NOTE: need different tool than the "embedding-files-into-app" example.
//
// Follow these steps first:
// $ go get -u github.com/kataras/bindata/cmd/bindata
// $ bindata ./assets/...
// $ go run main.go bindata_gzip.go
// "physical" files are not used, you can delete the "assets" folder and run the example.
// http://localhost:8080/ui/index.html

func newApp() *iris.Application {
	app := iris.New()

	/*
	Strange behavior of app.StaticEmbeddedGzip. It cannot detech and render index page automatically
	http://localhost:8080/index.html works
	but
	http://localhost:8080/ not found
	*/
	app.StaticEmbeddedGzip("/", "./assets/", GzipAsset, GzipAssetNames)

	/* I have to do trick to redirect request / to /index.html but showing index.html
	in browser address bar is really ugly
	 */

	/*
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/index.html", http.StatusMovedPermanently)
	})
	*/

	/*
	I propose another bad trick to write content of index.html when request is /
	Chrome, FireFox render correctly but Safari does not work at all
	 */
	app.Get("/", func (ctx iris.Context) {
		if data, err := GzipAsset("assets/index.html"); err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.WriteString("index.html is not found")
			return
		} else {
			ctx.StatusCode(http.StatusOK)
			ctx.Header("Content-Encoding", "gzip")
			ctx.ContentType("text/html")
			ctx.WriteGzip(data)
		}
	})
	/*
		Mean while app.StaticWeb can detect index page correctly
	 */
	//app.StaticWeb("/", "./assets/")

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}