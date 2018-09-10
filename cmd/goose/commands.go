package main

type gooseCommand interface {
	Execute([]string) error
	Name() string
	Desc() string
}

type dirOpts struct {
	Dir string `short:"d" long:"dir" description:"directory with migration files"`
}

type tableOpts struct {
	Table string `short:"t" long:"table" description:"the version Table"`
}

type migrationOpts struct {
	dirOpts
	tableOpts
}

type UpCommand struct{ migrationOpts }

func (cmd *UpCommand) Name() string {
	return "up"
}

func (cmd *UpCommand) Desc() string {
	return "Migrate the DB to the most recent version available"
}

func (cmd *UpCommand) Execute(args []string) error {
	panic("implement me")
}

type UpToCommand struct{ migrationOpts }

func (cmd *UpToCommand) Name() string {
	return "up-to"
}

func (cmd *UpToCommand) Desc() string {
	return "Migrate the DB to a specific VERSION"
}

func (cmd *UpToCommand) Execute(args []string) error {
	panic("implement me")
}

type DownCommand struct{ migrationOpts }

func (cmd *DownCommand) Name() string {
	return "down"
}

func (cmd *DownCommand) Desc() string {
	return "Roll back the version by 1"
}

func (cmd *DownCommand) Execute(args []string) error {
	panic("implement me")
}

type DownToCommand struct{ migrationOpts }

func (cmd *DownToCommand) Name() string {
	return "down-to"
}

func (cmd *DownToCommand) Desc() string {
	return "Roll back to a specific VERSION"
}

func (cmd *DownToCommand) Execute(args []string) error {
	panic("implement me")
}

type RedoCommand struct{ migrationOpts }

func (cmd *RedoCommand) Name() string {
	return "redo"
}

func (cmd *RedoCommand) Desc() string {
	return "Re-run the latest migration"
}

func (cmd *RedoCommand) Execute(args []string) error {
	panic("implement me")
}

type ResetCommand struct{ migrationOpts }

func (cmd *ResetCommand) Name() string {
	return "reset"
}

func (cmd *ResetCommand) Desc() string {
	return "Roll back all migrations"
}

func (cmd *ResetCommand) Execute(args []string) error {
	panic("implement me")
}

type CreateCommand struct{ dirOpts }

func (cmd *CreateCommand) Name() string {
	return "create"
}

func (cmd *CreateCommand) Desc() string {
	return "Creates new migration file with next version"
}

func (cmd *CreateCommand) Execute(args []string) error {
	panic("implement me")
}

type StatusCommand struct{ tableOpts }

func (cmd *StatusCommand) Name() string {
	return "status"
}

func (cmd *StatusCommand) Desc() string {
	return "Dump the migration status for the current DB"
}

func (cmd *StatusCommand) Execute(args []string) error {
	panic("implement me")
}

type VersionCommand struct{ tableOpts }

func (cmd *VersionCommand) Name() string {
	return "version"
}

func (cmd *VersionCommand) Desc() string {
	return "Print the current version of the database"
}

func (cmd *VersionCommand) Execute(args []string) error {
	panic("implement me")
}

/** register all commands to the parser */

func init() {
	var upCmd = new(UpCommand)
	var uptoCmd = new(UpToCommand)
	var downCmd = new(DownCommand)
	var downtoCmd = new(DownToCommand)
	var redoCmd = new(RedoCommand)
	var resetCmd = new(ResetCommand)
	var statusCmd = new(StatusCommand)
	var versionCmd = new(VersionCommand)
	var createCmd = new(CreateCommand)

	for _, cmd := range []gooseCommand{upCmd, uptoCmd, downCmd, downtoCmd, redoCmd, resetCmd, statusCmd, versionCmd, createCmd} {
		parser.AddCommand(cmd.Name(), cmd.Desc(), cmd.Desc(), cmd)
	}
}
