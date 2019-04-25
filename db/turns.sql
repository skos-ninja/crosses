create table if not exists turns
(
    id         text      not null,
    game_id    text      not null,
    player_id  text      not null,
    x_axis     integer   not null,
    y_axis     integer   not null,
    created_at timestamp not null,
    constraint turns_pk
        primary key (id),
    constraint turns_games_id_fk
        foreign key (game_id) references games
);

create index if not exists turns_game_id_index
    on turns (game_id);

create unique index if not exists turns_id_uindex
    on turns (id);
