// Combined boot server
package main

import (
	"flag"
	"fmt"
)

func main() {
	logFlag := flag.Int("v", 0, "logging level 0=critcal , 5=debug")
	flag.Parse()
	LogSetup(*logFlag)
	fmt.Println(banner)
	logger.Critical("Starting Astralboot Server")
	logger.Critical("Use -v=0-4 for extra logging")
	conf := GetConfig("config.toml")
	if *logFlag > 0 {
		logger.Critical("-- Implied Config Start --")
		conf.PrintConfig()
		logger.Critical("-- Implied Config Finish --")
	}
	// leases json database
	leases := NewStore(conf)
	logger.Info("starting tftp")
	go tftpServer(conf)
	logger.Info("start dhcp")
	go dhcpServer(conf, leases)
	logger.Info("start web server")
	wh := NewWebServer(conf, leases, *logFlag)
	go wh.Run()
	logger.Critical("Serving ...")
	// goroutine spinner
	c := make(chan int, 1)
	<-c
}

const banner = `
┏━┓┏━┓╺┳╸┏━┓┏━┓╻  ┏┓ ┏━┓┏━┓╺┳╸
┣━┫┗━┓ ┃ ┣┳┛┣━┫┃  ┣┻┓┃ ┃┃ ┃ ┃
╹ ╹┗━┛ ╹ ╹┗╸╹ ╹┗━╸┗━┛┗━┛┗━┛ ╹
https://github.com/zignig/astralboot
`
