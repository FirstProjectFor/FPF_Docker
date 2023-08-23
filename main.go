package main

import (
	"flag"
	"github.com/FirstProjectFor/FPF_Docker/logger"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strconv"
)

type Config struct {
	Server server
	Log    log
}

type server struct {
	Name string
	Port int
}

type log struct {
	File string
}

func main() {

	configFile := flag.String("c", "./configs/config_dev.toml", "-c configFile file")
	flag.Parse()

	config := &Config{}
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	panicIfNoNil(err)

	err = viper.Unmarshal(config)
	panicIfNoNil(err)

	logger.InitLogger(config.Log.File)

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		sugarLog := logger.GetLogger().Sugar()
		defer sugarLog.Sync()
		requestData, innerErr := io.ReadAll(req.Body)
		if innerErr != nil {
			sugarLog.Infof("read data failed, err:%v", err)
			return
		}
		sugarLog.Infof("accept reauest, request data:%s", string(requestData))

		//write data to response
		_, innerErr = res.Write([]byte("pong"))
		if innerErr != nil {
			sugarLog.Infof("write data failed, err:%v", err)
			return
		}
	})
	err = http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), nil)

	panicIfNoNil(err)
}

func panicIfNoNil(err error) {
	if err != nil {
		panic(err)
	}
}
