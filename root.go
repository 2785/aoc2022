package main

import (
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func c(e error) {
	if e != nil {
		panic(e)
	}
}

// nolint // go away
var l *zap.Logger
var s *zap.SugaredLogger

var start time.Time

var rootCmd = &cobra.Command{
	Use: "aoc2022",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		conf := zap.NewDevelopmentConfig()
		conf.DisableCaller = true
		conf.DisableStacktrace = true

		logger, err := conf.Build()
		c(err)

		l = logger
		s = logger.Sugar()

		start = time.Now()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		s.Infof("took %s", time.Since(start))
	},
}

func main() {
	c(rootCmd.Execute())
}
