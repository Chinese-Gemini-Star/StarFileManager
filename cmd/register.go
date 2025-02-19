package cmd

import (
	"StarFileManager/internal/call"
	"StarFileManager/internal/view"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// registerCmd 注册
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "用户注册",
	Long:  `进行用户登录,必须包含用户名,不能重复注册,允许后置输入密码`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}
		log.Debugln("username:", username, ", password:", password)

		if password == "" {
			i := view.NewPasswordInput(cmd.Context(), username, call.Register)
			p := tea.NewProgram(i)
			if _, err := p.Run(); err != nil {
				return err
			}
			return nil
		}
		err = call.Register(cmd.Context(), username, password)
		return err
	},
}

func init() {
	registerCmd.Flags().StringP("username", "u", "", "用户名,必须包含")
	registerCmd.MarkFlagRequired("username")
	registerCmd.Flags().StringP("password", "p", "", "密码,如果为空则后续手动输入")

	rootCmd.AddCommand(registerCmd)
}
