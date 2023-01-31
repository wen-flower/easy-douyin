package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wen-flower/easy-douyin/pkg/version"
)

var versionFlag bool

// NewCommand 创建一个 cobra.Command 对象
func NewCommand(name string, run func() error) *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名称，该名字会出现在帮助信息中
		Use:          name,
		SilenceUsage: true,
		// 执行 Command.Execute() 时执行的函数，函数会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			if versionFlag {
				fmt.Println(version.Get())
				return nil
			}
			return run()
		},
		// 检查参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q 不需要任何参数，得到 %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVarP(&versionFlag, "version", "v", false, "打印版本信息然后退出")

	return cmd
}
