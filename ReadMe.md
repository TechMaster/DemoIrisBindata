Demo [bindata](https://github.com/kataras/bindata) to embed static resources into iris web app.
**How to run this demo**

1. Run ```$ go get -u github.com/kataras/bindata/cmd/bindata```
2. Clone this repo ```$ git clone https://github.com/TechMaster/DemoIrisBindata```
3. Move into this repo ```cd DemoIrisBindata```
4. Run ```$ bindata ./assets/...``` to generate bindata_gzip.go. In repo, I already generated it.
5. Run ```$ go run main.go bindata_gzip.go```
6. Browse to http://localhost:8080/ you will see browser redirects to http://localhost:8080/index.html

The problem is
```app.StaticEmbeddedGzip("/", "./assets/", GzipAsset, GzipAssetNames)``` does not render default index.html
inside embedded asset folder

## Updated!

I work around by rendering index.html manually from embedded asset when app receive request "/"
Chrome, FireFox render correctly but Safary fails !
```go
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
    }})
```