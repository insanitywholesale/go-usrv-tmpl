package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go-usrv-tmpl/domain"
)

// TODO: implement Find and Store
// TODO: add "create table if not exists"
// TODO: check about using pgx instead of pq

type postgresRepo struct {
	client   *sql.DB
	database string
}

// TODO: integer might be wrong in the following sql
var createtable = `CREATE TABLE if not exists things (
	field TEXT,
	url TEXT,
	number INTEGER
);`

func newPostgresClient(postgresURL string) (*sql.DB, error) {
	client, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createtable)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresRepo(postgresURL string) (domain.ThingRepo, error) {
	repo := &postgresRepo{
		database: postgresURL,
	}
	client, err := newPostgresClient(postgresURL)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (r *postgresRepo) Find(field string) (*domain.Thing, error) {
	query := `SELECT * FROM pastes where field=$1`
	row := r.client.QueryRow(query, field)
	fmt.Println("row", row)
	var thing domain.Thing
	err := row.Scan(&thing.Field, &thing.URL, &thing.Number)
	if err != nil {
		return nil, err
	}
	return &thing, nil
}

func (r *postgresRepo) Store(thing *domain.Thing) error {
	query := `INSERT INTO things (field, url, number) VALUES ($1, $2, $3)`
	_, err := r.client.Exec(query, thing.Field, thing.URL, thing.Number)
	if err != nil {
		return err
	}
	return nil
}
