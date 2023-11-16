/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// 新增其他的cmd就是透過創造檔案的方式，cobra-cli add [fileName]
// 然後cmd的名稱就是 &cobra.Command{Use:[cmdName]}，要使用其他cmd就是go run main.go [cmdName]
// 如果是rootcmd就不用特別指明cmdName，直接操作go run main.go就好
// 而每個不同的cmd後面能加的flag也是獨立的，使用方法: go run main.go [cmdName] -[flagName Abbreviation]
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: "Randomly Gernerating A DADJOKE!!",
	Long:  `Getting Start and Have Your First Dadjoke!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: display, // 執行這個cmd會執行的函數
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dadjoke.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// 註冊flag，Bool代表的是該flag回傳值，也有Int等等
	// rootCmd.Flags().[returnType]P([flagName], [flagName Abbreviation], [defaultValue], [description])
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func display(cmd *cobra.Command, args []string) {
	fmt.Println("Hello World!")
}
