package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opstack",
	Short: "1-click OP Stack deployment CLI",
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the OP Stack to Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Deploying OP Stack...")
		exec.Command("make", "kind").Run()
		exec.Command("make", "deploy").Run()
	},
}

func Execute() {
	deployCmd.Flags().StringP("env", "e", "dev", "Environment to deploy")
	rootCmd.AddCommand(deployCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
