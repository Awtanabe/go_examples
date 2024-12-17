package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// ルートコマンド
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "This is a simple Cobra example",
		Long:  "This is a simple example to demonstrate Cobra CLI with commands and error handling.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command!")
		},
	}

	// サブコマンド: greet
	var greetCmd = &cobra.Command{
		Use:   "greet [name]",
		Short: "Greet someone",
		Args:  cobra.ExactArgs(1), // 引数は1つだけ必須
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	// サブコマンド: fail (エラーの例)
	var failCmd = &cobra.Command{
		Use:   "fail",
		Short: "An example of a command that fails",
		RunE: func(cmd *cobra.Command, args []string) error {
			// エラーを返す
			return errors.New("this command failed intentionally")
		},
	}

	// サブコマンドを追加
	rootCmd.AddCommand(greetCmd, failCmd)

	// コマンドの実行
	if err := rootCmd.Execute(); err != nil {
		// エラー発生時の処理
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

