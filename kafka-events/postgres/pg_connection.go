package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"time"
)

type DbService struct {
	host        string
	port        string
	schema      string
	username    string
	password    string
	connection  *sql.DB
	isConnected bool
}

func NewDbService(host string, port string, schema string, username string, password string) *DbService {
	return &DbService{
		host:        host,
		port:        port,
		schema:      schema,
		username:    username,
		password:    password,
		connection:  nil,
		isConnected: false,
	}
}

func (service *DbService) Connect(timeout time.Duration) error {
	if service.isConnected {
		fmt.Println("[WARN] sql: already connected to database")
		return nil
	}
	fmt.Println("[INFO] sql: connecting to database")
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		service.host, service.port, service.username, service.password, service.schema)

	connection, err := Connect("postgres", psqlInfo, timeout)
	if err == nil {
		fmt.Println("[INFO] sql: successfully connected to database")
		service.connection = connection
		service.isConnected = true
		return nil
	} else {
		return fmt.Errorf("[WARN] sql: failed connecting to database. \n\t%s", err)
	}
}

func (service *DbService) Disconnect() error {
	if service.isConnected == false || service.connection == nil {
		fmt.Println("[WARN] sql: already disconnected from database")
		return nil
	}
	err := service.connection.Close()
	if err != nil {
		fmt.Println("[INFO] sql: successfully disconnected from database")
		service.connection = nil
		service.isConnected = false
	} else {
		fmt.Println("[WARN] sql: failed disconnecting from database")
	}
	return err
}

func (service *DbService) Migrate() error {
	if service.isConnected == false {
		return fmt.Errorf("[ERROR] sql: migration cannot occur when disconnected from the database")
	}
	driver, err := postgres.WithInstance(service.connection, &postgres.Config{})
	if err != nil {
		return fmt.Errorf(
			"[ERROR] sql: migration cannot occur when disconnected from the database. \n\t%s",
			err)
	}
	instance, err := migrate.NewWithDatabaseInstance("file://migrations", service.schema, driver)
	if err != nil {
		return fmt.Errorf(
			"[ERROR] sql: migration failed. \n\t%s",
			err)
	}
	err = instance.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf(
			"[ERROR] sql: migration failed. \n\t%s",
			err)
	}
	return nil
}

// ConnectLoop tries to connect to the DB under given DSN using a give driver
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop.
func Connect(driver, DSN string, timeout time.Duration) (*sql.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			db, err := sql.Open(driver, DSN)
			if err == nil {
				return db, nil
			}
			log.Println(errors.Wrapf(err, "failed to connect to db %s", DSN))
		}
	}
}
