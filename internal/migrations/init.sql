create table if not exists tasks(
    id serial primary key,
    title text not null,
    content text,
    status boolean default false,
    created_at timestamp default current_timestamp
)