/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/lai0xn/lsm-cli/server"
	"github.com/lai0xn/lsm-cli/utils"
	"github.com/spf13/cobra"
)

var dir_path string

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "zip a directory and share it",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.ZipDir(dir_path)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		server.Serve("./out/output.zip", false)
	},
}

func init() {
	dirCmd.Flags().StringVarP(&dir_path, "path", "p", "", "dir path")
	dirCmd.MarkFlagRequired("dir")
	rootCmd.AddCommand(dirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
