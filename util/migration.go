package util

import (
	"database/sql"
	"strings"

	migrate "github.com/golang-migrate/migrate/v4"
	_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
)

type Migration struct {
	Migrate *migrate.Migrate
}

func (m *Migration) Up() (bool, error) {
	err := m.Migrate.Up()
	if err != nil && err != migrate.ErrNoChange {
		return false, err
	}
	return true, nil
}

func (m *Migration) Down() (bool, error) {
	err := m.Migrate.Down()
	if err != nil {
		return false, err
	}
	return true, err
}

func RunMigration(dbConn *sql.DB, migrationsFolderLocation string) (*Migration, error) {
	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationsFolderLocation)

	pathToMigrate := strings.Join(dataPath, "")

	driver, err := _postgres.WithInstance(dbConn, &_postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, MigrationDbName, driver)
	if err != nil {
		return nil, err
	}
	return &Migration{Migrate: m}, nil
}
