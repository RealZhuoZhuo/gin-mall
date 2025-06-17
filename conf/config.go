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
	Address   string `yaml:"address"`
	SmtpHost  string `yaml:"smtpHost"`
	SmtpEmail string `yaml:"smtpEmail"`
	SmtpPass  string `yaml:"smtpPass"`
}

type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
}

type Oss struct {
	BucketName      string `yaml:"BucketName"`
	QiNiuServer     string `yaml:"QiNiuServer"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
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
	if err := viper.Unmarshal(&Conf); err != nil { //**Config 类型,传递Conf的指针，Viper 会在内部初始化 Config 结构体，并将 Conf 指向这个新分配的内存
		log.Fatalf("序列化失败:%v", err)
	}
}
