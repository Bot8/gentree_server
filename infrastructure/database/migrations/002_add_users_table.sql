create schema general;

create table general.users
(
    id   serial not null
        constraint users_pk
            primary key,
    name text
);

create unique index user_id_uindex
    on general.users (id);

create table general.user_tree
(
    user_id integer not null
        constraint user_tree_users_id_fk
            references general.users
            on update cascade on delete cascade
            deferrable initially deferred,
    tree_id integer not null
        constraint user_tree_trees_fk
            references tree.trees
            on update cascade on delete cascade
            deferrable initially deferred,
    constraint user_tree_pk
        primary key (user_id, tree_id)
);

---- create above / drop below ----

drop table general.user_tree;
drop table general.users cascade;
drop schema general;