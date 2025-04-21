/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nnee2810/mimi-core/config"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
	"quizzy-classroom/model"
	"quizzy-classroom/server"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var serviceConfig model.ServiceConfig
		envPath, _ := cmd.Flags().GetString("env-path")

		if err := config.LoadConfig(envPath, &serviceConfig); err != nil {
			logger.Error("failed to load config", zap.Error(err))
			return
		}

		server.Run(&serviceConfig)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().StringP("env-path", "e", ".env", "Path to environment file")
}
