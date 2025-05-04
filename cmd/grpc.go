/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nnee2810/mimi-core/config"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
	"quizzy-classroom/grpc"
	"quizzy-classroom/model"

	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
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

		grpc.Run(&serviceConfig)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	grpcCmd.Flags().StringP("env-path", "e", ".env", "Path to environment file")
}
