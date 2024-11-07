package main

import (
	"github.com/spf13/cobra"

	"github.com/pkg/errors"
)

const (
	DefaultPostgresConStr = "tcp://127.0.0.1:26657"
	DefaultRedisConStr    = "127.0.0.1:9090"
)

func main() {
	rootCmd := &cobra.Command{
		Use:               "job test",
		Short:             "a demo for job test",
		PersistentPreRunE: nil,
	}

	rootCmd.AddCommand(
		StartCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		println(errors.Wrap(err, "rootCmd"))
	}

}
