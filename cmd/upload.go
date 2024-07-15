/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lai0xn/lsm-cli/server"
	"github.com/spf13/cobra"
)

var path string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "receive files from another device",
	Long:  "this command allows you to receive files from a different device",
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(path, true)
	},
}

func init() {
	uploadCmd.Flags().StringVarP(&path, "path", "p", "./myfiles/", "receive files")
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
