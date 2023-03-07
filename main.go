package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
)

var (
	localProxy    string
	clientDistDir string
	port          string
	vercelURL     string
)

func init() {
	flag.StringVar(&localProxy, "proxy", "", "proxy address, eg: http://localhost:1081")
	flag.StringVar(&clientDistDir, "client-dir", "dist/client", "dist/client directory path")
	flag.StringVar(&port, "port", ":9999", "server listen port")
	flag.StringVar(&vercelURL, "vercel-url", "https://chatgpt-demo-2972343.vercel.app", "like: https://chatgpt-demo-2972343.vercel.app")
}

func main() {
	flag.Parse()

	initSystem()

	g := gin.Default()
	g.Use(favicon.New(staticFileMap[ico])) // set favicon middleware
	g.GET("/", handleIndex)
	g.Any("/api/*path", handleApi)
	g.Any("/_astro/*path", handleStatic)
	g.GET("/favicon.svg", handleStatic)
	g.Run(port)
}

const (
	cssName    = "_astro/index"
	generateJs = "_astro/Generator"
	clientJs   = "_astro/client"
	webJs      = "_astro/web"
	ico        = "favicon.svg"
)

var (
	staticFileMap = map[string]string{} // name => path
	appURL        *url.URL
	htmlData      string
)

func initSystem() {
	err := filepath.Walk(clientDistDir, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, cssName) {
			staticFileMap[cssName] = path
		}
		if strings.Contains(path, generateJs) {
			staticFileMap[generateJs] = path
		}
		if strings.Contains(path, clientJs) {
			staticFileMap[clientJs] = path
		}
		if strings.Contains(path, webJs) {
			staticFileMap[webJs] = path
		}
		if strings.Contains(path, ico) {
			staticFileMap[ico] = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	appURL, err = url.Parse(vercelURL)
	if err != nil {
		panic(err)
	}

	downloadIndexFile()
}

func handleApi(ctx *gin.Context) {
	r := ctx.Request
	ctx.Request.Host = appURL.Host
	pxy := httputil.NewSingleHostReverseProxy(appURL)
	if localProxy != "" {
		u, _ := url.Parse(localProxy)
		pxy.Transport = &http.Transport{
			Proxy: http.ProxyURL(u),
		}
	}
	pxy.ServeHTTP(ctx.Writer, r)
}

func handleStatic(ctx *gin.Context) {
	for k, v := range staticFileMap {
		if strings.Contains(ctx.Request.RequestURI, k) {
			http.ServeFile(ctx.Writer, ctx.Request, v)
			return
		}
	}
	ctx.AbortWithStatus(404)
}

func handleIndex(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.String(200, htmlData)
}

func downloadIndexFile() {
	log.Println("[downloadIndexFile] start download index file")

	req, err := http.NewRequest(http.MethodGet, appURL.String(), nil)
	cli := http.Client{}
	if localProxy != "" {
		pxy, _ := url.Parse(localProxy)
		cli.Transport = &http.Transport{
			Proxy: http.ProxyURL(pxy),
		}

		log.Println("[downloadIndexFile] start download by proxy: ", pxy.String())
	}

	resp, err := cli.Do(req)
	if err != nil {
		panic("failed to download app index html: " + err.Error())
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("failed to download app index html: " + err.Error())
		return
	}
	resp.Body.Close()
	htmlData = string(data)

	log.Println("[downloadIndexFile] download index success")
}
