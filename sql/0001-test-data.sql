-- populate ingredients table
insert into ingredients (id, name, category) values ('123e4567-e89b-12d3-a456-426614174001', 'tomato', 'vegetable');
insert into ingredients (id, name, category) values ('123e4567-e89b-12d3-a456-426614174002', 'lettuce', 'vegetable');
insert into ingredients (id, name, category) values ('123e4567-e89b-12d3-a456-426614174003', 'cheese', 'protein');
insert into ingredients (id, name, category) values ('123e4567-e89b-12d3-a456-426614174004', 'ham', 'protein');
insert into ingredients (id, name, category) values ('123e4567-e89b-12d3-a456-426614174005', 'pita', 'bread');

-- populate orders table
insert into orders (id, user_id) values ('123e4567-e89b-12d3-a456-426614174000', '6ba7b814-9dad-11d1-80b4-00c04fd430c9');
insert into orders (id, user_id) values ('323e4567-e89b-12d3-a456-426614174000', '6ba7b814-9dad-11d1-80b4-00c04fd430c9');

-- populate orders_ingredients table
insert into orders_ingredients (order_id, ingredient_id) values ('123e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174001');
insert into orders_ingredients (order_id, ingredient_id) values ('123e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174005');
insert into orders_ingredients (order_id, ingredient_id) values ('123e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174004');

insert into orders_ingredients (order_id, ingredient_id) values ('323e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174001');
insert into orders_ingredients (order_id, ingredient_id) values ('323e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174002');
insert into orders_ingredients (order_id, ingredient_id) values ('323e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174003');
insert into orders_ingredients (order_id, ingredient_id) values ('323e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174004');
