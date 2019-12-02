create table employees (
    id varchar primary key,
    accountname varchar,
    displayname varchar,
    email varchar
);

create table folder (
    id varchar references employees(id),
    size varchar,
    time varchar
);