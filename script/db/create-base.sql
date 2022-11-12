create table visit (
                       id serial not null unique,
                       project_id bigint not null,
                       date date not null,
                       page_views bigint not null,
                       visitors bigint not null,

                       primary key (id)
);

create table project (
                         id serial not null unique,
                         name varchar(255) not null,
                         slug varchar(255),
                         currency varchar(10),
                         date_start date not null,
                         date_end date not null,
                         date_end_extra_time date,

                         primary key (id)
);

create table contribution (
                              id serial not null unique,
                              amount numeric not null,
                              user_id bigint not null,
                              project_id bigint not null,
                              created_at date not null,

                              primary key (id)
);

alter table visit add foreign key(project_id) references project;
alter table contribution add foreign key(project_id) references project;
