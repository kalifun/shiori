//go:generate go run assets-generator.go

package main

import (

	// Database driver

	_ "github.com/go-sql-driver/mysql"
	"github.com/kalifun/shiori/internal/cmd"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	// Add this to prevent it removed by go mod tidy
	_ "github.com/shurcooL/vfsgen"
)

func main() {
	err := cmd.ShioriCmd().Execute()
	if err != nil {
		logrus.Fatalln(err)
	}
}
