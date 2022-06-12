package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
	"os"
	"strings"
)

var (
	cfg  = flag.StringP("config", "c", "", "Configuration file.")
	help = flag.BoolP("help", "h", false, "Show this help message")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	// 使用环境变量。
	os.Setenv("KINGOFZIHUA_SERVER_GRPC_ADDR", "127.0.0.1:9000")

	viper.AutomaticEnv()                                              // 读取环境变量
	viper.SetEnvPrefix("KINGOFZIHUA")                                 // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))  // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
	viper.BindEnv("server.grpc.addr", "KINGOFZIHUA_SERVER_GRPC_ADDR") // 绑定环境变量名到key

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("json") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉 viper 以何种格式解析文件
	} else {
		viper.AddConfigPath(".")                  // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.kingofzihua") // 配置文件搜索路径，可以设置多个配置文件搜索路径
		viper.SetConfigName("config")             // 配置文件名称(没有文件扩展名) 在搜索时 Viper 会在文件名之后追加文件扩展名，并尝试搜索所有支持的扩展类型。
	}

	// 读取配置文件，如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file :%s \n", err))
	}

	fmt.Printf("Used configuration file is: %s \n", viper.ConfigFileUsed()) // 获取 config 文件

	fmt.Printf("read config server.http.addr is: %s \n", viper.Get("server.http.addr")) // 读取config参数

	// 反序列化为结构体
	var C Config

	err := viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct %v", err))
	}

	fmt.Printf("read config data.database.source is: %s \n", C.Data.Database.Source)

	// 序列化成字符串

	c := viper.AllSettings()

	bs, err := yaml.Marshal(c)

	if err != nil {
		panic(fmt.Sprintf("unable to mashal config to YAML: %v", err))
	}

	fmt.Printf("mashal config to YAML: \n %s \n", bs)
}

type Server struct {
	Http struct {
		Addr   string `mapstructure:"addr"`
		Timout string `mapstructure:"timout"`
	}
	Grpc struct {
		Addr   string `mapstructure:"addr"`
		Timout string `mapstructure:"timout"`
	}
}
type Data struct {
	Database struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	}
	Redis struct {
		Addr        string `mapstructure:"addr"`
		ReadTimeout string `mapstructure:"read_timeout"`
		WritTimeout string `mapstructure:"write_timeout"`
	}
}

type Config struct {
	Server Server `mapstructure:"server"`
	Data   Data   `mapstructure:"data"`
}
