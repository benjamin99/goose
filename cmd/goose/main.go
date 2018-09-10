package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	// Init DB drivers.
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/ziutek/mymysql/godrv"
)

var parser = flags.NewParser(nil, flags.Default|flags.None)

func main() {
	args, err := parser.Parse()
	if err != nil {
		log.Fatal("failed to parse args", err)
	}

	fmt.Println("args: ", args)
	//
	//if len(args) > 1 && args[0] == "create" {
	//	if err := goose.Run("create", nil, opt.Dir, opt.Table, args[1:]...); err != nil {
	//		log.Fatalf("goose run: %v", err)
	//	}
	//	return
	//}
	//
	//if len(args) < 4 {
	//	var b bytes.Buffer
	//
	//	parser.WriteHelp(&b)
	//	fmt.Println(b.String())
	//	return
	//}
	//
	//driver, dbstring, command := args[1], args[2], args[3]
	//
	//if err := goose.SetDialect(driver); err != nil {
	//	log.Fatal(err)
	//}
	//
	//switch driver {
	//case "redshift":
	//	driver = "postgres"
	//case "tidb":
	//	driver = "mysql"
	//}
	//
	//switch dbstring {
	//case "":
	//	log.Fatalf("-dbstring=%q not supported\n", dbstring)
	//default:
	//}
	//
	//db, err := sql.Open(driver, dbstring)
	//if err != nil {
	//	log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	//}
	//
	//var arguments []string
	//if len(args) > 3 {
	//	arguments = append(arguments, args[3:]...)
	//}
	//
	//if err := goose.Run(command, db, opt.Dir, opt.Table, arguments...); err != nil {
	//	log.Fatalf("goose run: %v", err)
	//}
}
