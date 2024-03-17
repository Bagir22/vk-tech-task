package Queries

const InsertUserQuery = `insert into "User" (name, balance) values ($1, $2);`
const InsertQuestQuery = `insert into "Quest" (name, cost) values ($1, $2);`
const InsertUserHistoryQuery = `insert into "UserHistory" (user_id, quest_id) values ($1, $2);`

const GetQuestCostQuery = `select cost from "Quest" where quest_id = $1`
const GetUserBalanceQuery = `select * from "User" where user_id = $1;`

const UpdateUserBalanceQuery = `update "User" set balance = $1 where user_id = $2;`
