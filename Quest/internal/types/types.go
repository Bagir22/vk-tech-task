package types

type User struct {
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
}

type UserFromDb struct {
	UserId int `db:"user_id"`
	User
}

type Quest struct {
	Name string `json:"name" db:"name"`
	Cost int    `json:"cost" json:"cost"`
}

type QuestFromDb struct {
	QuestId int `db:"quest_id"`
	Quest
}
type Response struct {
	Message     string `json:"message"`
	Description any    `json:"description"`
}

type Signal struct {
	UserId  int `json:"user_id"`
	QuestId int `json:"quest_id"`
}

type UserHistory struct {
	UserId    int    `db:"user_id" json:"user_id"`
	UserName  string `db:"username" json:"user_name"`
	QuestId   int    `db:"quest_id" json:"quest_id"`
	QuestName string `db:"quest_name" json:"quest_name"`
	Cost      int    `db:"cost" json:"cost"`
}
