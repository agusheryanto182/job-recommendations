package main

import (
	"cv-service/config"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
)

const dialect = "mysql"

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "./database/migrations", "directory with migration files")
)

// Database Migration CLI Command
//
// Goose Documentation: https://github.com/pressly/goose
// Reference: https://learning-cloud-native-go.github.io/docs/a8.adding_initial_database_migrations
func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}
	args = append(args, "-allow-missing")
	command := args[0]
	switch command {
	case "create":
		args := []string{args[1], "sql"}
		if err := goose.Run("create", nil, *dir, args...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}

	db, err := config.Initialize()
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()

	defer sqlDB.Close()
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatal(err)
	}

	switch command {
	case "up":
		if err := goose.Up(sqlDB, *dir, goose.WithAllowMissing()); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}
	if err := goose.Run(command, sqlDB, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
