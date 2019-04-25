create table if not exists games
(
    id             text not null,
    player1_id     text not null,
    player2_id     text,
    started_at     timestamp,
    finished_at    timestamp,
    winning_player text,
    constraint game_pk
        primary key (id)
);

create unique index if not exists game_id_uindex
    on games (id);