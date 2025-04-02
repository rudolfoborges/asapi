create table if not exists accounts (
    id uuid not null,
    name text not null,
    email text not null,
    password text not null,
    active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    primary key (id)
);
