package main

import "fmt"
import "os"

import "github.com/codegangsta/cli"
import "github.com/codegangsta/cli/altsrc"

func main() {

	// インスタンスのフィールドへ代入する都合上、サブサブなどの下位の方から作った方がうまくいく
	// main <main-flag> sub-command <sub-command-flag> sub-sub-command <sub-sub-command-flag> arguments
	subSubCommand := cli.Command{}
	subSubCommand.Name = "subsub"
	subSubCommand.Action = func(ctx *cli.Context) {
		opt := ctx.String("test-sub-sub")
		fmt.Println("cli-example sub subsub --test-sub-sub =", opt)
	}
	subSubCommand.Flags = []cli.Flag{
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:  "test-sub-sub",
				Value: "test-sub-sub default value",
			},
		),
		cli.StringFlag{
			Name:  "load",
			Value: "./.rc3",
		},
	}
	subSubCommand.Before = altsrc.InitInputSourceWithContext(
		subSubCommand.Flags,
		altsrc.NewYamlSourceFromFlagFunc("load"),
	)

	// サブコマンド
	// サブコマンドの下にも App と同じ設定を作ってみる
	// これはちゃんと動作するが、App とは別に設定しないとパースされない
	subCommand := cli.Command{}
	subCommand.Name = "sub"
	subCommand.Action = func(ctx *cli.Context) {
		opt := ctx.String("test-sub")
		fmt.Println("cli-example sub --test-sub =", opt)
	}
	subCommand.Flags = []cli.Flag{
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:  "test-sub",
				Value: "test-sub default value",
			},
		),
		cli.StringFlag{
			Name:  "load",
			Value: "./.rc2",
		},
	}
	subCommand.Before = altsrc.InitInputSourceWithContext(
		subCommand.Flags,
		altsrc.NewYamlSourceFromFlagFunc("load"),
	)
	// Command のサブコマンドは Subcommands で設定する
	subCommand.Subcommands = []cli.Command{
		subSubCommand,
	}

	// アプリケーション
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		altsrc.NewStringFlag(
			cli.StringFlag{
				Name:  "test",
				Value: "test default value",
			},
		),
		cli.StringFlag{
			Name:  "load",
			Value: "./.rc1",
		},
	}
	app.Action = func(ctx *cli.Context) {
		opt := ctx.String("test")
		fmt.Println("cli-example --test =", opt)
	}
	app.Before = altsrc.InitInputSourceWithContext(
		app.Flags,
		altsrc.NewYamlSourceFromFlagFunc("load"),
	)
	app.Commands = []cli.Command{
		subCommand,
	}
	app.Run(os.Args)
}
