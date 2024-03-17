package types

type User struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type UserFromDb struct {
	UserId  int    `db:"user_id"`
	Name    string `db:"name"`
	Balance int    `db:"balance"`
}

type Quest struct {
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

type QuestFromDb struct {
	QuestId int    `db:"quest_id"`
	Name    string `db:"name"`
	Cost    int    `db:"cost"`
}

type Response struct {
	Message     string `json:"message"`
	Description any    `json:"description"`
}
