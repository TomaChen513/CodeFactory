package lib

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

//server config datastructure
type ServerConfig struct {
	RunMode         string

	HTTPPort        int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration

	AppId           string
	RedirectURI     string
	AppKey          string
}

func LoadServerConfig() ServerConfig {
	var err error

	Cfg,err:=ini.Load("config/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}
	// load server
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	//load github config
	github, err := Cfg.GetSection("github")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	return ServerConfig{
		RunMode: Cfg.Section("").Key("RUN_MODE").MustString("debug"),
		HTTPPort:        server.Key("HTTP_PORT").MustInt(),
		ReadTimeout:     time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second,
		WriteTimeout:    time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second,
		AppId:           github.Key("APP_ID").MustString(""),
		AppKey:          github.Key("APP_KEY").MustString(""),
		RedirectURI:     github.Key("REDIRECT_URI").MustString(""),
	}

}