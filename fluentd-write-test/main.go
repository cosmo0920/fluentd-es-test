package main

import (
	"os"
	"strconv"
	"time"

	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	if len(os.Args) != 2 {
		log.Error("miss total param in command line")
		return
	}
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Error(err.Error())
		return
	}
	start := time.Now().UnixNano()
	var i int64
	for i = 0; i < n; i++ {
		log.Infof("n:%d i:%8d", n, i)
	}
	end := time.Now().UnixNano()
	var percent float64
	percent = float64(n) / (float64(end-start) / float64(time.Second))
	log.Infof(fmt.Sprintf("percent %.f/s", percent))
}
