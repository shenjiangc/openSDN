package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	logLevel      string
	logFile       string
	logMaxSize    int
	logMaxAge     int
	logMaxBackups int

	cqlCluster     string
	cqlPort        int
	cqlKeyspace    string
	cqlVersion     string
	cqlRF          int
	cqlConsistency string
	cqlRetry       int
	cqlCompress    string
	cqlTimeout     int

	rootCmd = &cobra.Command{
		Use:   "opensdn",
		Short: "open source sdn controller",
		Long:  `High performance SDN controller`,
	}
)

func initCmd() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	/*rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "log level")
	rootCmd.PersistentFlags().StringVar(&logFile, "logfile", "./log/opensdn.log", "log file path name")
	rootCmd.PersistentFlags().IntVar(&logMaxSize, "logsize", 100, "log max size(default 100MB)")
	rootCmd.PersistentFlags().IntVar(&logMaxAge, "logage", 30, "log max age(default 30days)")
	rootCmd.PersistentFlags().IntVar(&logMaxBackups, "lognum", 5, "log max backup num(default 5)")*/

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("opensdn")        // name of config file (without extension)
		viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("/etc/opensdn/")  // path to look for the config file in
		viper.AddConfigPath("$HOME/.opensdn") // call multiple times to add many search paths
		viper.AddConfigPath(".")              // optionally look for config in the working directory
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Printf("config file err, %s, %v\n", viper.ConfigFileUsed(), err)
		}
		os.Exit(2)
	} else {
		fmt.Printf("config file used, %s\n", viper.ConfigFileUsed())
	}

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "./log/opensdn.log")
	viper.SetDefault("log.size", 100)
	viper.SetDefault("log.age", 30)
	viper.SetDefault("log.num", 5)

	viper.SetDefault("cql.cluster", "127.0.0.1")
	viper.SetDefault("cql.port", 9042)
	viper.SetDefault("cql.keyspace", "opensdn")
	viper.SetDefault("cql.version", "3.0.0")
	viper.SetDefault("cql.replication_factor", 1)
	viper.SetDefault("cql.consistency", "quorum")
	viper.SetDefault("cql.retry", 3)
	viper.SetDefault("cql.compress", "")
	viper.SetDefault("cql.timeout", 600)

	logLevel = viper.GetString("log.level")
	logFile = viper.GetString("log.file")
	logMaxSize = viper.GetInt("log.size")
	logMaxAge = viper.GetInt("log.age")
	logMaxBackups = viper.GetInt("log.num")

	cqlCluster = viper.GetString("cql.cluster")
	cqlPort = viper.GetInt("cql.port")
	cqlKeyspace = viper.GetString("cql.keyspace")
	cqlVersion = viper.GetString("cql.version")
	cqlRF = viper.GetInt("cql.replication_factor")
	cqlConsistency = viper.GetString("cql.consistency")
	cqlRetry = viper.GetInt("cql.retry")
	cqlCompress = viper.GetString("cql.compress")
	cqlTimeout = viper.GetInt("cql.timeout")

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of OpenSDN",
	Long:  `All software has versions. This is OpenSDN's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
		os.Exit(0)
	},
}
