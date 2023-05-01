package db

import (
	"errors"
	"fmt"
	"go-grpc-services-course/internal/rocket"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

// New - returns a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return Store{}, err
	}

	return Store{db: db}, nil
}

// GetRocketById - retrieves a rocket from the db by id
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket

	row := s.db.QueryRow(
		`SELECT id, type, name FROM rockets where id=$1;`,
		id,
	)
	err := row.Scan(&rkt.ID, &rkt.Type, &rkt.Name)

	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - insert rocket in db
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		`INSERT INTO rockets (ID, NAME, TYPE)
		VALUES (:id, :name, :type)
		`,
		rkt,
	)

	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, errors.New("failed to insert in database")
	}

	return rocket.Rocket{
		ID:   rkt.ID,
		Type: rkt.Type,
		Name: rkt.Name,
	}, nil
}

// DeleteRocket - delete rocket from db
func (s Store) DeleteRocket(id string) error {
	// uid, err := uuid.from

	log.Println("inside2")

	_, err := s.db.DB.Exec(
		`DELETE FROM rockets where id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
