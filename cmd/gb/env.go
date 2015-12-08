package main

import (
	"github.com/constabulary/gb"
	"github.com/constabulary/gb/cmd"
)

func init() {
	registerCommand(EnvCmd)
}

var EnvCmd = &cmd.Command{
	Name:      "env",
	UsageLine: `env [var ...]`,
	Short:     "print project environment variables",
	Long: `
Env prints project environment variables. If one or more variable names is 
given as arguments, info prints the value of each named variable on its own line.
`,
	Run: info,
	ParseArgs: func(ctx *gb.Context, cwd string, args []string) []string {
		return args
	},
}
