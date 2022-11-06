// пользователи
CREATE TABLE users (
    id serial primary key,
    name varchar(255),
    username varchar(30),
    password_hash varchar(255)
);

// todo список
CREATE TABLE todo_list (
    id serial primary key,
    title varchar(255),
    description varchar(255)
);

// todo элементы списков(задачи)
CREATE TABLE todo_item (
    id serial primary key,
    title varchar(255),
    description varchar(255),
    done boolean
);

CREATE TABLE users_lists (
    id serial primary key,
    user_id integer references users(id),
    list_id integer references todo_list(id)
)

CREATE TABLE list_items (
    id serial primary key,
    list_id integer references todo_list(id)
    item_id integer references todo_item(id)
)