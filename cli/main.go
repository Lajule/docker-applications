package main

import (
	"github.com/Lajule/docker-applications"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "docker-applications",
		Short: "docker-applications - Run docker-compose commands on multiple apps",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var config docker_applications.Config

			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Err:%s", err)
			}

			if err := docker_applications.Execute(args, config); err != nil {
				log.Fatalf("Err:%s", err)
			}
		},
	}
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("docker-applications")
	}

	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Err:%s", err)
	}
}

func init() {
	gotenv.Load()

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "config file (default is ./docker-applications.yml)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Err:%s", err)
	}
}
