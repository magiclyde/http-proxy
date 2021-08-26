/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/9 下午5:37
 * @note:
 */

package internal

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port     int    // listen port
	Proto    string // listen protocol
	CertFile string // path to certificate file
	KeyFile  string // path to certificate key file
}

func NewConfig() *Config {
	v := viper.New()
	v.SetDefault("port", 8888)
	v.SetDefault("proto", "https")
	v.SetDefault("certFile", "./assets/tls/server.pem")
	v.SetDefault("keyFile", "./assets/tls/server.key")
	v.AddConfigPath("./etc/")
	v.AddConfigPath("/etc/")
	v.SetConfigName("http-proxy")
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Read config err: %v, use default settings", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unmarshal config err: %+v", err)
	}

	return &cfg
}
