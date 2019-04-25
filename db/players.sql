create table if not exists players
(
    id         text      not null,
    name       text      not null,
    created_at timestamp not null,
    constraint players_pk
        primary key (id)
);

create unique index if not exists players_id_uindex
    on players (id);
