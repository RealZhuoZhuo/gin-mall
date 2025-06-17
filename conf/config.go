package conf

import (
	"github.com/spf13/viper"
	"log"
)

var Conf *Config

type Config struct {
	System        System        `yaml:"system"`
	Mysql         Mysql         `yaml:"mysql"`
	Email         Email         `yaml:"email"`
	EncryptSecret EncryptSecret `yaml:"encryptSecret"`
	Oss           Oss           `yaml:"oss"`
}

type Mysql struct {
	Dns string `yaml:"dns"`
}

type Email struct {
	Address   string      `yaml:"address"`
	SmtpHost  interface{} `yaml:"smtpHost"`
	SmtpEmail interface{} `yaml:"smtpEmail"`
	SmtpPass  interface{} `yaml:"smtpPass"`
}

type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
}

type Oss struct {
	BucketName      interface{} `yaml:"BucketName"`
	QiNiuServer     interface{} `yaml:"QiNiuServer"`
	AccessKeyId     interface{} `yaml:"AccessKeyId"`
	AccessKeySecret interface{} `yaml:"AccessKeySecret"`
}

type System struct {
	Env    string `yaml:"env"`
	Port   string `yaml:"port"`
	Domain string `yaml:"domain"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读入配置文件出错:%v", err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("序列化失败:%v", err)
	}

}
