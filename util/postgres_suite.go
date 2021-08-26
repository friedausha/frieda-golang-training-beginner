package util

import (
	"database/sql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	// Register pq-timeouts driver to database/sql
	_ "github.com/Kount/pq-timeouts"

	"github.com/stretchr/testify/suite"
)

const (
	MigrationDbName = "postgres"
	PostgresDriver  = "pq-timeouts"

	// DefaultTestDsn is the default url for testing postgresql in the postgres test suites
	DefaultTestDsn = "user=frieda password=namamu dbname=golang_training host=localhost port=5432 sslmode=disable read_timeout=60000 write_timeout=60000"
)

// Suite struct for MySQL Suite
type Suite struct {
	suite.Suite
	DSN                     string
	DBConn                  *sql.DB
	Migration               *Migration
	MigrationLocationFolder string
	DBName                  string
}

// SetupSuite setup at the beginning of test
func (s *Suite) SetupSuite() {
	var err error
	s.DBConn, err = sql.Open(PostgresDriver, s.DSN)
	s.Require().NoError(err)
	err = s.DBConn.Ping()
	s.Require().NoError(err)
	s.Migration, err = RunMigration(s.DBConn, s.MigrationLocationFolder)
	s.Require().NoError(err)
}

// TearDownSuite teardown at the end of test
func (s *Suite) TearDownSuite() {
	err := s.DBConn.Close()
	s.Require().NoError(err)
}
