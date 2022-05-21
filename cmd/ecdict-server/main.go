package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	_ "github.com/youguanxinqing/log-conf"
	"os"
	v1 "telegram-bot/api/v1"
	"telegram-bot/internal/config"
)

var (
	port = flag.Int("port", 5121, "listen port")
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
	r.GET("/", v1.EcdictHandle)

	if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
		log.Fatalf("run panic: %v", err)
	}
}
