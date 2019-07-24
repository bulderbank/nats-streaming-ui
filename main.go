package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Masterminds/sprig"
	"github.com/bulderbank/nats-streaming-ui/models"
	"github.com/bulderbank/nats-streaming-ui/utils"
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
		c.Redirect(302, "/channels")
	})

	router.GET("/channels", func(c *gin.Context) {
		url := "http://localhost:8222/streaming/channelsz?subs=1"

		chs := models.NatsChannels{}
		err := utils.JsonGet(url, &chs)
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "nats.html", gin.H{
			"title": "NATS Streaming Channels",
			"nats":  chs,
		})
	})

	router.GET("/channels/:channel", func(c *gin.Context) {
		url := "http://localhost:8222/streaming/channelsz?channel=" + c.Param("channel") + "&subs=1"

		ch := models.NatsChannel{}
		err := utils.JsonGet(url, &ch)
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "channel.html", gin.H{
			"title":   fmt.Sprintf("NATS Streaming Channel: %s", ch.Name),
			"channel": ch,
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
