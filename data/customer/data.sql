create table if not exists customers (
    id bigint primary key,
    name varchar(255) not null,
    age int
);

insert into customers(id, name, age) values(1, 'Harry', 27);
insert into customers(id, name, age) values(2, 'Ginny', 22);
insert into customers(id, name, age) values(3, 'Ron', 27);
insert into customers(id, name, age) values(4, 'Hermoine', 27);
insert into customers(id, name, age) values(5, 'Lupin', 48);
insert into customers(id, name, age) values(6, 'Hagrid', 89);
