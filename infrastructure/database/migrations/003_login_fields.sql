alter table general.users
    add login text not null;

alter table general.users
    add password text not null;

create unique index users_login_uindex
    on general.users (login);

---- create above / drop below ----

alter table general.users
    drop column login;

alter table general.users
    drop column password;
