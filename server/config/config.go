package config

import "github.com/spf13/viper"

var LoadConf *Conf

type Conf struct {
	Viper *viper.Viper
}

func loadConfigFromYaml(c *Conf) error {
	c.Viper = viper.New()

	//设置配置文件的名字
	c.Viper.SetConfigName("server")

	//添加配置文件所在的路径
	c.Viper.AddConfigPath("./etc")

	//设置配置文件类型
	c.Viper.SetConfigType("yml")

	if err := c.Viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func init() {
	LoadConf = new(Conf)
	loadConfigFromYaml(LoadConf)
}
