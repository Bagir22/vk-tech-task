package postgres

import (
	"Quest/Database/Queries"
	"Quest/internal/config"
	"Quest/internal/types"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func InitDb(db *sqlx.DB) *Db {
	return &Db{
		db: db,
	}
}

func InitConn() (*sqlx.DB, error) {
	cfg := config.InitConfig()

	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"db", cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDatabase, "disable")

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Db) AddUser(user types.User) error {
	_, err := d.db.Exec(Queries.InsertUserQuery, user.Name, user.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) AddQuest(quest types.Quest) error {
	_, err := d.db.Exec(Queries.InsertQuestQuery, quest.Name, quest.Cost)
	if err != nil {
		return err
	}
	return nil
}
