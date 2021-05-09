package MySQL

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Query(host, userName, pass string, port int, dbName, pemPath, format, input string) error {
	if err := OpenDB(host, userName, pass, port, dbName, pemPath); err != nil {
		return err
	}
	defer db.Close()

	buf, err := os.ReadFile(input)
	if err != nil {
		return err
	}
	s := string(buf)

	rows, err := db.Query(s)
	if err != nil {
		return nil
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([]sql.RawBytes, len(cols))

	if format == "TSV" {
		fmt.Println(strings.Join(cols, "\t"))
	} else {
		fmt.Println(strings.Join(cols, ","))
	}

	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		value := []string{}
		for _, col := range vals {
			if col == nil {
				value = append(value, "")
			} else {
				value = append(value, string(col))
			}
		}
		if format == "TSV" {
			fmt.Println(strings.Join(value, "\t"))
		} else {
			fmt.Println(strings.Join(value, ","))
		}
	}

	return nil

}

func registerTlsConfig(pemPath, tlsConfigKey string) error {
	caCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(pemPath)
	if err != nil {
		return err
	}

	if ok := caCertPool.AppendCertsFromPEM(pem); !ok {
		return fmt.Errorf("Failed to append pem")
	}

	mysql.RegisterTLSConfig(tlsConfigKey, &tls.Config{
		ClientCAs:          caCertPool,
		InsecureSkipVerify: true,
	})

	return nil
}

func OpenDB(host, userName, pass string, port int, dbName, pemPath string) error {
	var err error
	if pemPath != "" {
		if err := registerTlsConfig(pemPath, "custom"); err != nil {
			return err
		}

		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=custom", userName, pass, host, port, dbName))
		if err != nil {
			return err
		}
	} else {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, pass, host, port, dbName))
		if err != nil {
			return err
		}
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(60 * time.Second)

	return nil
}
