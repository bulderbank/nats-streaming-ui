package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(sprig.HtmlFuncMap())
	router.HTMLRender = loadTemplates("./templates")

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	router.GET("/", func(c *gin.Context) {
		url := "http://localhost:8222/streaming/channelsz?subs=1"
		spaceClient := http.Client{
			Timeout: time.Second * 2, // Maximum of 2 secs
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "spacecount-tutorial")

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		chs := natsChannels{}
		jsonErr := json.Unmarshal(body, &chs)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		c.HTML(http.StatusOK, "nats.html", gin.H{
			"title": "NATS Streaming Queues",
			"nats":  chs,
		})
	})

	router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), sprig.HtmlFuncMap(), files...)
	}
	return r
}
