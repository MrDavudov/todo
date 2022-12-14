CREATE TABLE users (
    id              serial          not null unique,
    name            varchar(255)    not null,
    username        varchar(30)     not null,
    password_hash   varchar(255)    not null
);

CREATE TABLE todo_list (
    id          serial          not null unique,
    title       varchar(255)    not null,
    description varchar(255)    not null
);

CREATE TABLE todo_item (
    id          serial          not null unique,
    title       varchar(255)    not null,
    description varchar(255)    not null,
    done        boolean         not null default false
);

CREATE TABLE users_lists (
    id      serial  not null unique,
    user_id int     references users(id) on delete cascade      not null,
    list_id int     references todo_list(id) on delete cascade  not null
);

CREATE TABLE list_items (
    id      serial  not null unique,
    list_id int     references todo_list(id) on delete cascade  not null,
    item_id int     references todo_item(id) on delete cascade  not null
);