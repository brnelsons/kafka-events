CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table sources (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    name VARCHAR(36) UNIQUE NOT NULL,
    description VARCHAR
);

insert into sources (name, description) VALUES ('common', 'events that are available to all sources');

create table event_types (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    source_uuid UUID NOT NULL,
    name VARCHAR(36) NOT NULL,
    description VARCHAR,
    UNIQUE(source_uuid, name)
);

insert into event_types (source_uuid, name, description) VALUES
((select uuid from sources where name = 'common'), 'COMMENT', null);

create table events (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    type_uuid UUID NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    start_date_time timestamp NOT NULL,
    end_date_time timestamp NOT NULL,
    FOREIGN KEY (type_uuid) REFERENCES event_types(uuid)
);

create table event_associations (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    event_uuid UUID NOT NULL,
    source_uuid UUID NOT NULL,
    date_time TIMESTAMP NOT NULL,
    status VARCHAR,
    UNIQUE(event_uuid, source_uuid, date_time),
    FOREIGN KEY (source_uuid) REFERENCES sources(uuid),
    FOREIGN KEY (event_uuid) REFERENCES events(uuid)
);