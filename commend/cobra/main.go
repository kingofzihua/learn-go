package main

import (
	"fmt"
	"github.com/kingofzihua/learn-go/commend/cobra/internal/version"
	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:   "skonline",
	Short: "skonline Short message",
	Long:  `skonline long message `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	pflag.StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.skonline.yaml)")
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		h, err := home.Dir()
		if err != nil {
			panic(fmt.Errorf("can not found home work :%w", err))
		}
		viper.AddConfigPath(h)
		viper.SetConfigName(".skonline")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("can't read config: %w", err))
	}

	fmt.Printf("load config: %v", viper.AllSettings())
}
