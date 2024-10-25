package config

import (
	"fmt"
	"testing"
)

func TestConf(t *testing.T) {
	cfg := LoadConf
	fmt.Println(cfg.Viper.GetInt("redis.db"))
}
