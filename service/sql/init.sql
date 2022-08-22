create table orders (
    id text primary key,
    data jsonb
);

create table messages (
    subId text primary key,
    lastMsgId bigint
)