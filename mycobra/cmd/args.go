package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var cusArgsCheckCmd = &cobra.Command{
	Use: "cusArgs",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cusArgs cmd begin")
		fmt.Println(args)
		fmt.Println("run cusArgs cmd end")
	},
}

var argsCheckCmd = &cobra.Command{
	Use:       "args",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"123", "abc", "nick"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run cusArgs cmd begin")
		fmt.Println(args)
		fmt.Println("run cusArgs cmd end")
	},
}

func init() {
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsCheckCmd)
}
