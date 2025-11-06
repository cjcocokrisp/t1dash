package db

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/cjcocokrisp/t1dash/internal/config"
	"github.com/cjcocokrisp/t1dash/pkg/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

//go:embed setup-db.sh
var setupScript []byte

//go:embed migrations/*.sql
var MigrationFS embed.FS

// URL for the DB must be inited before using
var (
	DBUrl  string
	DBPool *pgxpool.Pool
)

func InitDBURL(hostname string, port int, database string, user string, password string) {
	DBUrl = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, hostname, port, database)
}

// InitDBConnection sets up the database connection
// If the database does not exist it creates it
// It also checks if the tables in the db exists and then creates those as well
func InitDBConnection() {
	if DBUrl == "" {
		log.WithFields(log.Fields{
			"error": "DBUrl not set when attempting to make connection!",
		}).Error("FATAL")
		os.Exit(1)
	}

	var err error
	DBPool, err = pgxpool.New(context.Background(), DBUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = DBPool.Ping(context.Background())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "28000" || pgErr.Code == "3D000" {
				log.Info("Database or user not created, beginning setup")
				setupDB()

				err = DBPool.Ping(context.Background())
				if err != nil {
					util.LogError("Failed to connect to database after creation", "setup", err)
					os.Exit(1)
				}
			} else {
				util.LogPgError(pgErr.Code, pgErr.Message)
				os.Exit(1)
			}
		} else { // TODO: add connection timeout error so it's not as messy
			util.LogError("An unhandled database error occured", "setup", err)
			os.Exit(1)
		}
	}
	log.Info("Database connection successful")

	// Check if tables exist and if they don't create them
	tables := []string{"users", "sessions"} // More tables will be added

	for _, t := range tables {
		var exists bool
		query := "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)"
		err = DBPool.QueryRow(context.Background(), query, t).Scan(&exists)
		if err != nil {
			util.LogError("Failed to query information_schema for table names", "setup", err)
			os.Exit(1)
		}

		if !exists {
			log.Info("A table was missing and needs to be added, running migrations")
			createTables()
			break
		}
	}
}

func CloseDBConnection() {
	DBPool.Close()
}

// setupDB creates the database and requested user
// Runs the shell script ~/scripts/setup-db.sh
func setupDB() {
	temp, err := os.CreateTemp("", "setup-db-*.sh")
	if err != nil {
		util.LogError("Creating temp file for database creation failed", "setup(db)", err)
		os.Exit(1)
	}
	defer os.Remove(temp.Name())

	_, err = temp.Write(setupScript)
	if err != nil {
		util.LogError("Writing setup script to temp file failed", "setup(db)", err)
		os.Exit(1)
	}
	temp.Close()

	err = os.Chmod(temp.Name(), 0755)
	if err != nil {
		util.LogError("Chmod for temp file failed", "setup(db)", err)
		os.Exit(1)
	}

	cmd := exec.Command(temp.Name(), config.AppCfg.DBHostname, strconv.Itoa(config.AppCfg.DBPort),
		config.AppCfg.DBRootPassword, config.AppCfg.DBUser, config.AppCfg.DBPassword,
		config.AppCfg.DBDatabase)
	out, err := cmd.CombinedOutput()
	if err != nil {
		util.LogError("Running setup script failed", "setup", err)
		log.WithFields(log.Fields{
			"output": string(out),
		}).Error("Command run output")
		os.Exit(1)
	}
	log.Info("Database created successfully")
}

// createTables runs the migrations for the project and creates the tables
func createTables() {
	d, err := iofs.New(MigrationFS, "migrations")
	if err != nil {
		util.LogError("Error setting up iofs for migrations", "setup(tables)", err)
		os.Exit(1)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, DBUrl+"?sslmode=disable")
	if err != nil {
		util.LogError("Error creating instance connected to DB", "setup(tables)", err)
		os.Exit(1)
	}

	err = m.Up()
	if err != nil {
		util.LogError("Error running migrations", "setup(db)", err)
		os.Exit(1)
	}
	log.Info("Migrations finished, successfully created tables")
}
