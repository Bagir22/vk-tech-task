package Queries

const InsertUserQuery = `insert into "User" (name, balance) values ($1, $2);`
const InsertQuestQuery = `insert into "Quest" (name, cost) values ($1, $2);`
const InsertUserHistoryQuery = `insert into "UserHistory" (user_id, quest_id) values ($1, $2);`

const GetQuestCostQuery = `select cost from "Quest" where quest_id = $1`
const GetUserQuery = `select * from "User" where user_id = $1;`

const UpdateUserBalanceQuery = `update "User" set balance = $1 where user_id = $2;`

const GetUserHistoryQuery = `select "User".user_id as "user_id", "User".name as "username", "Quest"."quest_id" as "quest_id",
        							"Quest"."name" as "quest_name", "Quest".cost from "UserHistory" 
        					    join "User" on "UserHistory".user_id = "User".user_id
                            	join "Quest" on "UserHistory".quest_id = "Quest".quest_id
                            	where "User".user_id = $1;`

const CheckUserExistQuery = `select count(*) from "User" where user_id = $1;`
