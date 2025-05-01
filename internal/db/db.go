package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ensomnatt/ducks/internal/config"
	"github.com/ensomnatt/ducks/internal/logger"
	"github.com/ensomnatt/ducks/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DucksDB struct {
	db *pgxpool.Pool
}

func NewDucksDB(db *pgxpool.Pool) *DucksDB {
	return &DucksDB{db: db}
}

func ConnectToDB() (*pgxpool.Pool, error) {
	config := config.GetConfig()
	logger.Log.Debug("got config", "config", config)
	pgConn := fmt.Sprintf(
		"postgres://%s:%s@db:5432/ducks",
		config.PostgresUser,
		config.PostgresPassword,
	)
	logger.Log.Debug("created pg connection variable", "connection", pgConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	logger.Log.Debug("created context")

	pool, err := pgxpool.New(ctx, pgConn)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to db: %v", err)
	}
	logger.Log.Debug("connected to db")

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("db isn't responsing: %v", err)
	}
	logger.Log.Debug("checked connection")
	return pool, nil
}

func (db *DucksDB) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS ducks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`

	_, err := db.db.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error while creating table: %v", err)
	}

	return nil
}

func (db *DucksDB) Create(duck models.Duck, ctx context.Context) error {
	query := "INSERT INTO ducks (name, age) VALUES ($1, $2)"
	_, err := db.db.Exec(ctx, query, duck.Name, duck.Age)

	logger.Log.Debug("creating a duck", "name", duck.Name, "age", duck.Age)
	
	if err != nil {
		return fmt.Errorf("error while creating duck: %v", err)
	}

	return nil
}

func (db *DucksDB) Get(name string, ctx context.Context) (models.Duck, error) {
	query := "SELECT (name, age) FROM ducks WHERE name = $1"
	var duck models.Duck
	err := db.db.QueryRow(ctx, query, name).Scan(&duck)

	logger.Log.Debug("getting a duck", "name", duck.Name, "age", duck.Age)

	if err != nil {
		return duck, fmt.Errorf("error while getting a duck: %v", err)
	}

	return duck, nil
}

func (db *DucksDB) GetAll(ctx context.Context) ([]models.Duck, error) {
	query := "SELECT name, age FROM ducks"
	rows, err := db.db.Query(ctx, query)
	
	if err != nil {
		return nil, fmt.Errorf("error while getting ducks: %v", err) 
	}

	defer rows.Close()

	var ducks []models.Duck

	for rows.Next() {
		var duck models.Duck
		err := rows.Scan(&duck.Name, &duck.Age)
		if err != nil {
			return nil, fmt.Errorf("error while scanning a row: %v", err)
		}

		ducks = append(ducks, duck)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("there are some errors while getting ducks: %v", rows.Err())
	}

	return ducks, nil
}
