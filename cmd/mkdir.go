package cmd

import (
	"StarFileManager/internal/call"
	"errors"
	"github.com/spf13/cobra"
)

// mkdirCmd 创建目录
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "创建目录",
	Long:  `在指定目录下创建目录,`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, _ := cmd.Flags().GetBool("parents")
		if len(args) < 1 {
			return errors.New("缺少操作数")
		}
		_, err := call.MakeDir(cmd.Context(), args[0], p)
		return err
	},
}

func init() {
	mkdirCmd.Flags().BoolP("parents", "p", false, "允许必要时自动创建父目录")

	rootCmd.AddCommand(mkdirCmd)
}
