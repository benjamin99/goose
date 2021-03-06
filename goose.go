package goose

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
)

var (
	duplicateCheckOnce sync.Once
	minVersion         = int64(0)
	maxVersion         = int64((1 << 63) - 1)
)

// Run runs a goose command.
func Run(command string, db *sql.DB, dir, table string, args ...string) error {

	fmt.Println("run: ", command, "dir: ", dir, "table: ", table)

	switch command {
	case "up":
		SetTableName(table)
		if err := Up(db, dir); err != nil {
			return err
		}
	case "up-by-one":
		SetTableName(table)
		if err := UpByOne(db, dir); err != nil {
			return err
		}
	case "up-to":
		if len(args) == 0 {
			return fmt.Errorf("up-to must be of form: goose [OPTIONS] DRIVER DBSTRING up-to VERSION")
		}

		version, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("version must be a number (got '%s')", args[0])
		}

		SetTableName(table)
		if err := UpTo(db, dir, version); err != nil {
			return err
		}
	case "create":
		if len(args) == 0 {
			return fmt.Errorf("create must be of form: goose [OPTIONS] DRIVER DBSTRING create NAME [go|sql]")
		}

		migrationType := "go"
		if len(args) == 2 {
			migrationType = args[1]
		}
		if err := Create(db, dir, args[0], migrationType); err != nil {
			return err
		}
	case "down":
		SetTableName(table)
		if err := Down(db, dir); err != nil {
			return err
		}
	case "down-to":
		if len(args) == 0 {
			return fmt.Errorf("down-to must be of form: goose [OPTIONS] DRIVER DBSTRING down-to VERSION")
		}

		version, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("version must be a number (got '%s')", args[0])
		}

		SetTableName(table)
		if err := DownTo(db, dir, version); err != nil {
			return err
		}
	case "redo":
		SetTableName(table)
		if err := Redo(db, dir); err != nil {
			return err
		}
	case "reset":
		SetTableName(table)
		if err := Reset(db, dir); err != nil {
			return err
		}
	case "status":
		SetTableName(table)
		if err := Status(db, dir); err != nil {
			return err
		}
	case "version":
		SetTableName(table)
		if err := Version(db, dir); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%q: no such command", command)
	}
	return nil
}
