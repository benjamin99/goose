package main

import (
	"database/sql"
	"github.com/benjamin99/goose"
	"github.com/jessevdk/go-flags"
	"log"
	"os"

	// Init DB drivers.
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	var opt struct {
		Dir   string `short:"d" long:"Dir" description:"directory with migration files"`
		Table string `short:"t" long:"Table" description:"the version Table"`
		Help  bool   `shot:"h" long:"Help"`
	}

	args, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatal("failed to parse args", err)
	}

	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, opt.Dir, opt.Table, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	if len(args) < 3 || opt.Help {
		usage()
		return
	}

	driver, dbstring, command := args[1], args[2], args[3]

	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	switch driver {
	case "redshift":
		driver = "postgres"
	case "tidb":
		driver = "mysql"
	}

	switch dbstring {
	case "":
		log.Fatalf("-dbstring=%q not supported\n", dbstring)
	default:
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, opt.Dir, opt.Table, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

func usage() {
	log.Print(usagePrefix)
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND

Drivers:
    postgres
    mysql
    sqlite3
    redshift

Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up

    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
    goose tidb "user:password@/dbname?parseTime=true" status

Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)
