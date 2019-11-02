package main

import (
	"flag"
	"fmt"
	zimg "github.com/hzde0128/gimg"
	_ "github.com/bradfitz/gomemcache/memcache"
	"net/http"
	"os"
	"runtime"
)

var cfgFile string
var zContext *zimg.ZContext

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPtr := flag.String("config", "", "config file")
	flag.Usage = usage
	flag.Parse()

	if *configPtr == "" {
		*configPtr = "./conf/config.ini"
	}

	cfgFile = *configPtr

	isExist, _ := exists(cfgFile)
	if !isExist {
		fmt.Println("config file not exist!")
		os.Exit(-1)
	}

	zContext, err := zimg.NewContext(cfgFile)
	checkError(err)
	defer zContext.Release()

	//zContext.Logger.Info("load config.ini success!")

	addr := fmt.Sprintf("%s:%d", zContext.Config.System.Host, zContext.Config.System.Port)
	zContext.Logger.Info("server start run :  %s", addr)

	zHttpd := zimg.NewHttpd(zContext)
	err = http.ListenAndServe(addr, zHttpd)
	if err != nil {
		zContext.Logger.Error("error : %s", err.Error())
	}

	//signalHandle()
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:--config=/etc/config.ini \n")
	flag.PrintDefaults()
	os.Exit(-2)
}

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(-2)
	}
}
