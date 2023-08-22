package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strconv"
)

type Config struct {
	Server server
}

type server struct {
	Name string
	Port int
}

func main() {
	configFile := flag.String("c", "./configs/config_dev.toml", "-c configFile file")
	flag.Parse()

	fmt.Println(*configFile)

	config := &Config{}

	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	panicIfNoNil(err)

	err = viper.Unmarshal(config)
	panicIfNoNil(err)

	fmt.Printf("config: %#v \n", config)

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		requestData, innerErr := io.ReadAll(req.Body)
		if innerErr != nil {
			fmt.Printf("read data failed, err:%s. \n", err)
			return
		}
		fmt.Println(string(requestData))

		//write data to response
		_, innerErr = res.Write([]byte("pong " + *configFile))
		if innerErr != nil {
			fmt.Printf("write data failed, err:%s. \n", err)
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
