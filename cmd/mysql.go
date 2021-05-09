package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yakawa/sql2csv/MySQL"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL",
	RunE:  MySQLRunE,
}

var (
	pem string
)

func init() {
	mysqlCmd.PersistentFlags().StringVar(&host, "host", "localhost", "hostname")
	mysqlCmd.PersistentFlags().IntVar(&port, "port", 3306, "port")
	mysqlCmd.PersistentFlags().StringVar(&db, "db", "", "Database")
	mysqlCmd.PersistentFlags().StringVar(&pem, "pem", "", "SSL Cert Path")
	mysqlCmd.PersistentFlags().StringVar(&user, "user", "", "Database")
	mysqlCmd.PersistentFlags().StringVar(&pass, "pass", "", "SSL Cert Path")
}

func MySQLRunE(cmd *cobra.Command, args []string) error {
	MySQL.Query(host, user, pass, port, db, pem, outputFormat, inputFile)
	return nil
}
