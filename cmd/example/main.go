package main

import (
	"flag"
	"github.com/exiaohao/golang-template/pkg/example/server"
	"github.com/spf13/cobra"
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
			server := new(server.HttpServer)
			stopCh := setupSignalHandler()

			server.Initialize(opts)
			server.Run(stopCh)
			return nil
		},
	}
)

func init() {
	exampleCmd.PersistentFlags().StringVar(&opts.Address, "address" , "0.0.0.0", "Server listen address, default 0.0.0.0")
	exampleCmd.PersistentFlags().Uint16Var(&opts.Port, "port", 9000, "Server listen port, default 9000")
	// optional, only for kubeClient
	// with default value, cobra will connect with default env config.
	exampleCmd.PersistentFlags().StringVar(&opts.KubeConfig, "kubeconfig", "", "path to kubeconfig file")
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
