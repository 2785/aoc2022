package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func c(e error) {
	if e != nil {
		panic(e)
	}
}

var l *zap.Logger
var s *zap.SugaredLogger

var rootCmd = &cobra.Command{
	Use: "aoc2022",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewDevelopment()
		c(err)

		l = logger
		s = logger.Sugar()
	},
}

func main() {
	c(rootCmd.Execute())
}
