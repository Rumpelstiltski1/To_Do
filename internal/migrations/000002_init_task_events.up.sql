create table if not exist task_event(
    id serial primary key,
    task_id int not null,
    action TEXT not null,
    title TEXT,
    content TEXT,
    status boolean,
    created_at TIMESTAMP NOT NULL DEFAULT now()
    )