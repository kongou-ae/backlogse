// Copyright © 2017 kongou-ae
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	//	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "backlose",
	Short: "close the issue of backlog",
	Long:  "close the issue of backlog",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: set id of issue.")
			os.Exit(-1)
		}

		if viper.GetString("url") == "" {
			fmt.Println("Error: set url in .backlose.")
			os.Exit(-1)
		}

		if viper.GetString("apikey") == "" {
			fmt.Println("Error: set apikey in .backlose.")
			os.Exit(-1)
		}

		if viper.GetString("projectname") == "" {
			fmt.Println("Error: set projectname in .backlose.")
			os.Exit(-1)
		}

		backlogUrl := viper.GetString("url")
		req, err := http.NewRequest(
			"PATCH",
			backlogUrl+"/api/v2/issues/"+viper.GetString("projectname")+"-"+args[0]+"?apiKey="+viper.GetString("apikey")+"&statusId=4",
			nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		if resp.StatusCode == 200 {
			fmt.Println(viper.GetString("projectname") + "-" + args[0] + " is closed.")
		} else {
			fmt.Println("Error: " + viper.GetString("projectname") + "-" + args[0] + " is not closed")
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is default $HOME/.backlose.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().StringP("apikey", "", "", "apikye of backlog")
	RootCmd.PersistentFlags().StringP("projectname", "", "", "projectname of backlog")
	RootCmd.PersistentFlags().StringP("url", "", "", "url")

	viper.BindPFlag("apikey", RootCmd.PersistentFlags().Lookup("apikey"))
	viper.BindPFlag("projectname", RootCmd.PersistentFlags().Lookup("projectname"))
	viper.BindPFlag("uel", RootCmd.PersistentFlags().Lookup("url"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".backlose") // name of config file (without extension)
	viper.AddConfigPath(".")         // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error: config file don't exist")
		os.Exit(-1)
	}

}
