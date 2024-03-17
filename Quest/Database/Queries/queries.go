package Queries

const InsertUserQuery = `insert into "User" (name, balance) values ($1, $2);`
const InsertQuestQuery = `insert into "Quest" (name, cost) values ($1, $2);`
