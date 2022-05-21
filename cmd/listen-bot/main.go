package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"os"
	v1 "telegram-bot/api/v1"
	"telegram-bot/internal/config"
	"telegram-bot/pkg/decorator/httputil"
)

var (
	port = flag.Int("port", 5120, "listen port")
	conf = flag.String("conf", "config.toml", "config path")
)

func init() {
	flag.Parse()

	bytes, err := os.ReadFile(*conf)
	log.Error(err)
	if err := toml.Unmarshal(bytes, &config.Basic); err != nil {
		log.Fatalf("%v", err)
	} else {
		log.Println(config.Basic)
	}
}

func main() {
	r := gin.Default()

	// v1
	{
		g := r.Group("/v1")
		g.GET("/", v1.HelloWorld)
		g.POST("/", v1.HelloWorld)
		g.POST("/:token", httputil.RequireToken(config.Basic.UrlToken, v1.BotHandle))
	}

	if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
		log.Fatalf("run panic: %v", err)
	}
}
