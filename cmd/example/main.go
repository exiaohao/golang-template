package main

import (
	"flag"
	"github.com/exiaohao/golang-template/pkg/example/server"
	"github.com/spf13/cobra"
	"k8s.io/apiserver/pkg/util/logs"
	"os"
	"os/signal"
	"syscall"
)

var (
	opts    server.InitOptions
	rootCmd = &cobra.Command{
		Use:          "example server name",
		Short:        "your short descr here",
		Long:         "your long descr here",
		SilenceUsage: true,
	}

	exampleCmd = &cobra.Command{
		Use:   "run",
		Short: "your short descr here",
		Long:  "your long descr here",
		RunE: func(*cobra.Command, []string) error {
			logs.InitLogs()
			defer logs.FlushLogs()

			server := new(server.HttpServer)
			stopCh := setupSignalHandler()

			server.Initialize(opts)
			server.Run(stopCh)
			return nil
		},
	}
)

func init() {
	exampleCmd.PersistentFlags().StringVar(&opts.Address, "Server listen address, default 0.0.0.0", "0.0.0.0", "")
	exampleCmd.PersistentFlags().Uint16Var(&opts.Port, "Server listen port, default 9000", 9000, "")
}

func main() {
	flag.Parse()

	rootCmd.AddCommand(exampleCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// setupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(stop)
		<-sigs
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
