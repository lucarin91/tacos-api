drop table if exists orders;
drop table if exists orders_ingredients;
drop table if exists ingredients;

create table orders (
    id char(36) not null,
    user_id char(36) not null,
    primary key (id, user_id)
);

create table orders_ingredients (
    order_id char(36) not null,
    ingredient_id char(36) not null
);

create table ingredients (
    id char(36) not null,
    name text not null,
    category text not null
);