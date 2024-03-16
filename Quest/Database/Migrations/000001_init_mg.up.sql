create schema quests;

create table if not exists quests."User"
(
    user_id serial  primary key,
    name    varchar(50) not null,
    balance integer     not null
);

create table if not exists quests."Quest"
(
    quest_id serial  primary key,
    name    varchar(50) not null,
    cost integer        not null
);

create table if not exists quests."UserHistory"
(
    user_id integer
        constraint UserHistory_user_id_fk
            references quests."User"
        on update restrict on delete restrict,
    quest_id integer
        constraint UserHistory_quest_id_fk
            references quests."Quest"
        on update restrict on delete restrict
);