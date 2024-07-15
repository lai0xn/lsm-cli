/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/lai0xn/lsm-cli/server"
	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var (
	file_path string
	shareCmd  = &cobra.Command{
		Use:   "file",
		Short: "share file in lan",
		Long:  `share is a command used to share a single file to another device acroos the network`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if the file exists
			if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
				fmt.Println("File Does not exist")
				return
			}
			// if the file exists serve the file
			server.Serve(file_path, false)
		},
	}
)

func init() {
	// add the required file flag and bind it to the file_path variable
	shareCmd.Flags().StringVarP(&file_path, "path", "p", "", "file path")
	shareCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(shareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
