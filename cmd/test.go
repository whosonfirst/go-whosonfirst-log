package main

import (
       "io"
       "os"
       "github.com/whosonfirst/go-whosonfirst-log"
)

func main() {

	writer := io.MultiWriter(os.Stdout)

	logger := log.NewWOFLogger("[your-app] ")
	logger.AddLogger(writer, "debug")

	logger.Info("Writing all your logs to %s", "wub wub wub")
	logger.Debug("Hello world")

	logger.Status("STATUS")

}
