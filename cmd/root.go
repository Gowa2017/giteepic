/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"giteepic/lib"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "giteepic {file1 [file2 file3 ... flieN]}",
	Short: "用来上传图片到 gitee 仓库",
	Long: `所有的需要参数由命令行进行完整指定，调用 gitee 的 OpenAPI 接口进行
	使用 giteepic -h 来看如何使用`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("args:%s", strings.Join(args, ","))
		lib.Upload(user, repo, token, path, args)
	},
	Args: cobra.MinimumNArgs(1),
}

var (
	token string
	user  string
	repo  string
	path  string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.giteepic.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&token, "token", "k", "", "gitee 生成的个人令牌")
	rootCmd.Flags().StringVarP(&user, "user", "u", "", "gitee用户名")
	rootCmd.Flags().StringVarP(&repo, "repo", "r", "", "gitee仓库名")
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "gitee仓库中用来存储图片的路径，默认是 /")
	rootCmd.MarkFlagRequired("token")
	rootCmd.MarkFlagRequired("user")
	rootCmd.MarkFlagRequired("repo")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".giteepic" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".giteepic")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
