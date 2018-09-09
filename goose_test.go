package goose

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDefaultBinary(t *testing.T) {
	commands := []string{
		"go build -i -o goose ./cmd/goose",
		"./goose -d examples/sql-migrations sqlite3 sql.db up",
		"./goose -d examples/sql-migrations sqlite3 sql.db version",
		"./goose -d examples/sql-migrations sqlite3 sql.db down",
		"./goose -d examples/sql-migrations sqlite3 sql.db status",
	}

	for _, cmd := range commands {
		args := strings.Split(cmd, " ")
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			t.Fatalf("%s:\n%v\n\n%s", err, cmd, out)
		}
	}
}

func TestDefaultBinaryWithSpecifiedTableName(t *testing.T) {
	commands := []string{
		"go build -i -o goose ./cmd/goose",
		"./goose -d examples/sql-migrations sqlite3 sql2.db up -t goose_db_version_2",
		"./goose -d examples/sql-migrations sqlite3 sql2.db version -t goose_db_version_2",
		"./goose -d examples/sql-migrations sqlite3 sql2.db down -t goose_db_version_2",
		"./goose -d examples/sql-migrations sqlite3 sql2.db status -t goose_db_version_2",
	}

	for _, cmd := range commands {
		args := strings.Split(cmd, " ")
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			t.Fatalf("%s:\n%v\n\n%s", err, cmd, out)
		}
	}
}



func TestCustomBinary(t *testing.T) {
	commands := []string{
		"go build -i -o custom-goose ./examples/go-migrations",
		"./custom-goose -dir=examples/go-migrations sqlite3 go.db up",
		"./custom-goose -dir=examples/go-migrations sqlite3 go.db version",
		"./custom-goose -dir=examples/go-migrations sqlite3 go.db down",
		"./custom-goose -dir=examples/go-migrations sqlite3 go.db status",
	}

	for _, cmd := range commands {
		args := strings.Split(cmd, " ")
		out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			t.Fatalf("%s:\n%v\n\n%s", err, cmd, out)
		}
	}
}
