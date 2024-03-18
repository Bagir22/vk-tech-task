package postgres

import (
	"Quest/Database/Queries"
	"Quest/internal/config"
	"Quest/internal/types"
	"context"
	"database/sql"
	"errors"
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

func (d *Db) AddUser(ctx context.Context, user types.User) error {
	_, err := d.db.Exec(Queries.InsertUserQuery, user.Name, user.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) AddQuest(ctx context.Context, quest types.Quest) error {
	_, err := d.db.Exec(Queries.InsertQuestQuery, quest.Name, quest.Cost)
	if err != nil {
		return err
	}
	return nil
}

func (d *Db) ProcessSignal(ctx context.Context, signal types.Signal) (types.User, error) {
	fail := func(err error) error {
		return fmt.Errorf("Process signal: %v", err)
	}

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return types.User{}, fail(err)
	}

	defer tx.Rollback()

	//Get price for the quest
	cost := 0
	if err = tx.QueryRowContext(ctx, Queries.GetQuestCostQuery, signal.QuestId).Scan(&cost); err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fail(fmt.Errorf("Quest is not exist"))
		}
		return types.User{}, fail(err)
	}
	//Get user
	var user types.UserFromDb
	if err = tx.QueryRowContext(ctx, Queries.GetUserQuery, signal.UserId).Scan(&user.UserId,
		&user.Name, &user.Balance); err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fail(fmt.Errorf("User is not exist"))
		}
		return types.User{}, fail(err)
	}

	//Marking the task completed
	_, err = tx.ExecContext(ctx, Queries.InsertUserHistoryQuery, signal.UserId, signal.QuestId)
	if err != nil {
		return types.User{}, fail(err)
	}

	//Update user balance
	newBalance := user.Balance + cost
	_, err = tx.ExecContext(ctx, Queries.UpdateUserBalanceQuery, newBalance, signal.UserId)
	if err != nil {
		return types.User{}, fail(err)
	}

	if err = tx.Commit(); err != nil {
		return types.User{}, fail(err)
	}

	return types.User{user.Name, user.Balance + cost}, nil
}

func (d *Db) GetUserHistory(ctx context.Context, id int) ([]types.UserHistory, error) {
	var count int
	err := d.db.QueryRow(Queries.CheckUserExistQuery, id).Scan(&count)
	if err != nil {
		return []types.UserHistory{}, err
	}
	if count == 0 {
		return []types.UserHistory{}, errors.New("User is not exist")
	}

	var history []types.UserHistory
	rows, err := d.db.Query(Queries.GetUserHistoryQuery, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var completedTask types.UserHistory
		err := rows.Scan(&completedTask.UserId, &completedTask.UserName,
			&completedTask.QuestId, &completedTask.QuestName, &completedTask.Cost)
		if err != nil {
			return nil, err
		} else {
			history = append(history, completedTask)
		}
	}

	return history, nil
}

func (d *Db) GetUsers(ctx context.Context) ([]types.UserFromDb, error) {
	var users []types.UserFromDb

	rows, err := d.db.Query(Queries.GetUsersQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user types.UserFromDb
		err := rows.Scan(&user.UserId, &user.Name, &user.Balance)
		if err != nil {
			return nil, err
		} else {
			users = append(users, user)
		}
	}

	return users, nil
}

func (d *Db) GetQuestById(ctx context.Context, id int) (types.QuestFromDb, error) {
	var quest types.QuestFromDb

	err := d.db.Get(&quest, Queries.GetQuestQuery, id)
	if err != nil {
		return types.QuestFromDb{}, err
	}

	return quest, nil
}

func (d *Db) UpdateQuest(ctx context.Context, quest types.Quest, id int) error {
	_, err := d.db.Exec(Queries.UpdateQuestQuery, quest.Name, quest.Cost, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Db) GetQuests(ctx context.Context) ([]types.QuestFromDb, error) {
	var quests []types.QuestFromDb

	rows, err := d.db.Query(Queries.GetQuestsQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var quest types.QuestFromDb
		err := rows.Scan(&quest.QuestId, &quest.Name, &quest.Cost)
		if err != nil {
			return nil, err
		} else {
			quests = append(quests, quest)
		}
	}

	return quests, nil
}
