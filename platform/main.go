package main

import (
	"platform/placeholder"
	"platform/services"
)

// func writeMessage(logger logging.Logger, cfg config.Configuration) {
// 	section, ok := cfg.GetSection("main")
// 	if ok {
// 		message, ok := section.GetString("message")
// 		if ok {
// 			logger.Info(message)
// 		} else {
// 			logger.Panic("Cannot find configuration setting")
// 		}
// 	} else {
// 		logger.Panic("Config section not found")
// 	}
// }

func main() {
	services.RegisterDefaultServices()
	placeholder.Start()

	// var cfg config.Configuration
	// services.GetService(&cfg)
	//
	// var logger logging.Logger
	// services.GetService(&logger)

	// writeMessage(logger, cfg)
	// services.Call(writeMessage)
	//
	// val := struct {
	// 	message string
	// 	logging.Logger
	// }{
	// 	message: "Hello from the struct",
	// }
	// services.Populate(&val)
	// val.Logger.Debug(val.message)
}
