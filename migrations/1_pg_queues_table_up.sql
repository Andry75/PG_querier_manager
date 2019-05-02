create table pg_queues
(
    id serial,
    ip varchar(16) not null,
    active_jobs_count int default 0 not null,
    not_active_since timestamp not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

create unique index pg_queues_id_uindex
    on pg_queues (id);

alter table pg_queues
    add constraint pg_queues_pk
        primary key (id);
