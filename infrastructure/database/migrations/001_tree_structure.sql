create schema tree;

create table tree.persons
(
    id          serial                              not null
        constraint person_pk
            primary key,
    first_name  text                                not null,
    last_name   text                                not null,
    middle_name text,
    birth       date,
    death       date,
    photo       text,
    history     text,
    created     timestamp default CURRENT_TIMESTAMP not null,
    updated     timestamp default CURRENT_TIMESTAMP not null
);

create unique index person_id_uindex
    on tree.persons (id);

create table tree.trees
(
    id      serial                              not null
        constraint tree_pk
            primary key,
    name    text                                not null,
    created timestamp default CURRENT_TIMESTAMP not null,
    updated timestamp default CURRENT_TIMESTAMP not null
);

create unique index tree_id_uindex
    on tree.trees (id);

create table tree.families
(
    id               serial                              not null
        constraint pair_pk
            primary key,
    first_parent_id  integer                             not null
        constraint families_persons_first_parent_person_fk
            references tree.persons
            on update cascade on delete restrict
            deferrable initially deferred,
    second_parent_id integer                             not null
        constraint families_persons_second_parent_person_fk
            references tree.persons
            on update cascade on delete restrict
            deferrable initially deferred,
    marriage         date,
    created          timestamp default CURRENT_TIMESTAMP not null,
    updated          timestamp default CURRENT_TIMESTAMP not null
);


create unique index family_id_uindex
    on tree.families (id);

create table tree.family_tree
(
    tree_id   integer not null
        constraint family_tree_trees_fk
            references tree.trees
            on update cascade on delete cascade
            deferrable initially deferred,
    family_id integer not null
        constraint family_tree_families_fk
            references tree.families
            on update cascade on delete cascade
            deferrable initially deferred,
    constraint family_tree_pk
        primary key (tree_id, family_id)
);


create table tree.family_child
(
    family_id integer not null
        constraint family_child_families_fk
            references tree.families
            on update cascade on delete cascade
            deferrable initially deferred,
    child_id  integer not null
        constraint family_child_persons_id_fk
            references tree.persons
            on update cascade on delete cascade
            deferrable initially deferred,
    constraint family_child_pk
        primary key (family_id, child_id)
);

---- create above / drop below ----

drop table tree.family_child;
drop table tree.family_tree;
drop table tree.families cascade;
drop table tree.trees cascade;
drop table tree.persons cascade;
drop schema tree;
