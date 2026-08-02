//go:build integration

package external

import (
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Requires a real MySQL instance - see ../docker-compose.yml (make entity-db-up) or
// override EMI_ENTITY_MYSQL_DSN. Run with: make entity-migration-test-mysql
//
// NOTE: this project targets Postgres primarily. uniqueId's column default
// (gorm:"type:uuid;default:gen_random_uuid()", see
// lib/golang/go-entity-default-fields.go) is Postgres-specific - MySQL has no
// equivalent built-in function, so AutoMigrate is expected to fail here with an
// unrecognized-function error. Left in place (rather than deleted) so a future MySQL
// port has something to start from; not something actively being kept green.
func TestAutoMigrate_MySQL(t *testing.T) {
	dsn := os.Getenv("EMI_ENTITY_MYSQL_DSN")
	if dsn == "" {
		dsn = "emi:emi@tcp(localhost:53306)/emi_entity_test?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		t.Fatalf("failed to connect to mysql (is `make entity-db-up` running? dsn=%q): %v", dsn, err)
	}

	runAutoMigrateAssertions(t, db)
}
