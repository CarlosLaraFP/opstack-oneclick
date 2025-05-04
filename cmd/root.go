package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
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
		chainId, _ := cmd.Flags().GetInt("chain-id")
		rpcUrl, _ := cmd.Flags().GetString("rpc-url")
		namespace, _ := cmd.Flags().GetString("namespace")

		log.Println("Creating namespace...")
		if err := exec.Command("kubectl", "create", "namespace", namespace).Run(); err != nil {
			log.Println(err)
		}

		log.Println("🔐 Creating JWT secret...")
		jwt := generateJWTSecret() // returns hex string
		const jwtPath = "/tmp/jwt.txt"
		if err := os.WriteFile(jwtPath, []byte(jwt), 0600); err != nil {
			log.Println(err)
		}

		// Run kubectl create secret
		secretCmd := exec.Command("kubectl", "create", "secret", "generic", "jwt-secret",
			"--from-file=jwt.txt="+jwtPath,
			"-n", namespace)
		if err := secretCmd.Run(); err != nil {
			log.Fatalf("❌ Failed to create secret: %v", err)
		}

		log.Println("🚀 Deploying OP Stack...")

		c := exec.Command("helm", "upgrade", "--install", "opstack", "./helm/opstack",
			"-n", namespace,
			"--set", fmt.Sprintf("opNode.chainId=%d", chainId),
			"--set", fmt.Sprintf("opNode.rpcUrl=%s", rpcUrl),
		)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			log.Fatalf("🚨 Helm error: %v", err)
		}

	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show pod status in the given namespace",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, err := cmd.Flags().GetString("namespace")
		if err != nil {
			log.Println(err)
		}

		fmt.Println("📦 Checking deployment status...")
		out, err := exec.Command("kubectl", "get", "pods", "-n", namespace).CombinedOutput()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		log.Println(string(out))
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the OP Stack deployment",
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")

		fmt.Println("Deleting Helm release...")
		if err := exec.Command("helm", "uninstall", "opstack", "-n", namespace).Run(); err != nil {
			log.Println(err)
		}

		fmt.Println("Deleting namespace...")
		if err := exec.Command("kubectl", "delete", "namespace", namespace).Run(); err != nil {
			log.Println(err)
		}
	},
}

// In a production OP Stack, rpc-url is how op-node connects to its execution engine (op-geth), exposing JSON-RPC on 8545.
func Init() {
	deployCmd.Flags().Int("chain-id", 420, "Chain ID for the OP Chain")
	deployCmd.Flags().String("rpc-url", "http://localhost:8545", "RPC URL")
	deployCmd.Flags().String("namespace", "opstack", "Kubernetes namespace for deployment")
	statusCmd.Flags().String("namespace", "opstack", "Kubernetes namespace")
	destroyCmd.Flags().String("namespace", "opstack", "Kubernetes namespace")

	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(destroyCmd)
}

func Execute() {
	Init()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func generateJWTSecret() string {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		log.Fatalf("Failed to generate JWT secret: %v", err)
	}
	return hex.EncodeToString(secret)
}
