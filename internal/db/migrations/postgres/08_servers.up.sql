begin;

-- For now at least the IDs will be the same as the name, because this allows us
-- to not have to persist some generated ID to worker and controller nodes.
-- Eventually we may want them to diverge, so we have both here for now.

create table servers (
    private_id text,
    type text,
    name text not null unique
      check(length(trim(name)) > 0),
    description text,
    address text,
    create_time wt_timestamp,
    update_time wt_timestamp,
    primary key (private_id, type)
  );
  
create trigger 
  immutable_columns
before
update on servers
  for each row execute procedure immutable_columns('create_time');
  
create trigger 
  default_create_time_column
before
insert on servers
  for each row execute procedure default_create_time();

commit;