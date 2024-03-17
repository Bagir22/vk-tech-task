create table if not exists "User"
(
    user_id serial  primary key,
    name    varchar(50) not null,
    balance integer     not null
    );

create table if not exists "Quest"
(
    quest_id serial  primary key,
    name    varchar(50) not null,
    cost integer        not null
    );

create table if not exists "UserHistory"
(
    user_id integer
    constraint UserHistory_user_id_fk
    references "User"
    on update restrict on delete restrict,
    quest_id integer
    constraint UserHistory_quest_id_fk
    references "Quest"
    on update restrict on delete restrict
);