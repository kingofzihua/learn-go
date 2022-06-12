package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var firstname string
var (
	lastname = pflag.StringP("lastname", "n", "zh", "Input Your Last Name")
	count    = pflag.Int("count", 0, "Count Number")
	version  = pflag.String("version", "v0.0.1", "Version")
	debug    = pflag.Bool("debug", false, "Is Debug")
)

// 上面的函数命名是有规则的：
//		函数名带 Var 说明是将标志的值绑定到变量，否则是将标志的值存储在指针中。
//		函数名带 P 说明支持短选项，否则不支持短选项。

func main() {
	pflag.StringVar(&firstname, "firstname", "wang", "Input Your First Name")

	// 设置 没有指定选项值时的默认值
	pflag.Lookup("version").NoOptDefVal = "v1.0.0"
	pflag.Lookup("lastname").NoOptDefVal = "wangzh"
	// version :
	// 		[nothing]    		=> v0.0.1
	// 		--version    		=> v1.0.0
	// 		--version=v1.0.1  	=>  v2.0.0

	// 开启一个新的 FlagSet
	var help string
	lagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	lagSet.StringVar(&help, "help", "help", "help usage")

	// 弃用指定命令
	// 弃用标志/速记会在帮助文本中隐藏它，并在使用弃用标志/速记时打印使用消息。
	pflag.CommandLine.MarkDeprecated("firstname", "firstname 已不被推荐使用")

	// 弃用短名
	pflag.CommandLine.MarkShorthandDeprecated("lastname", "请使用 --lastname")

	pflag.Parse() // 必须得在所有定义完成后执行

	fmt.Printf("argument number is :%v\n", pflag.NArg())   // 非选项参数的数量
	fmt.Printf("argument list is: %v\n", pflag.Args())     // 非选项参数
	fmt.Printf("the first argument is:%v\n", pflag.Arg(0)) // 获取第一个非选项参数
	fmt.Printf("firstname is :%v\n", firstname)
	fmt.Printf("lastname is :%v\n", *lastname)
	fmt.Printf("count is :%v\n", *count)
	fmt.Printf("version is :%v\n", *version)
	fmt.Printf("debug is :%v\n", *debug)
	first, _ := pflag.CommandLine.GetString("firstname")
	last, _ := pflag.CommandLine.GetString("lastname")
	fmt.Printf("get name is :%v %v\n", first, last)

	fmt.Printf("help is :%v \n", help)
}
