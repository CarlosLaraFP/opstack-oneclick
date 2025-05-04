package cmd

import (
	"bytes"
	"testing"
)

func TestDeployCommand(t *testing.T) {
	var buf bytes.Buffer
	Init()
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"deploy", "--chain-id", "999", "--rpc-url", "http://localhost:8545", "--namespace", "testns"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestStatusCommandHelp(t *testing.T) {
	rootCmd.SetArgs([]string{"status", "--help"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
