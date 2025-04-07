/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nnee2810/mimi-core/config"
	"github.com/nnee2810/mimi-core/logger"
	"go.uber.org/zap"
	"quizzy-classroom/migration"
	"quizzy-classroom/model"

	"github.com/spf13/cobra"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := logger.Init(); err != nil {
			panic(err)
		}

		var serviceConfig model.ServiceConfig
		envPath, _ := cmd.Flags().GetString("env-path")

		if err := config.LoadConfig(envPath, &serviceConfig); err != nil {
			logger.Error("failed to load config", zap.Error(err))
			return
		}

		migration.Run(&serviceConfig)
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	migrationCmd.Flags().StringP("env-path", "e", ".env", "Path to environment file")
}
