//go:build integration

package external

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Requires a real Postgres instance - see ../docker-compose.yml (make entity-db-up) or
// override EMI_ENTITY_POSTGRES_DSN. Run with: make entity-migration-test-postgres
func TestAutoMigrate_Postgres(t *testing.T) {
	dsn := os.Getenv("EMI_ENTITY_POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost port=55432 user=emi password=emi dbname=emi_entity_test sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		t.Fatalf("failed to connect to postgres (is `make entity-db-up` running? dsn=%q): %v", dsn, err)
	}

	runAutoMigrateAssertions(t, db)
}
