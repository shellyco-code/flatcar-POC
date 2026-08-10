-- +migrate Up

-- 1. Drop redundant duplicate index on instance_application (exact duplicate of primary key (instance_id, application_id))
drop index if exists instance_application_instance_id_application_id_idx;

-- 2. Drop redundant single-column index on instance_application(instance_id) (covered by PK leading column)
drop index if exists instance_application_instance_id_idx;

-- 3. Drop redundant single-column index on instance_application(group_id) (covered by composite group_id_last_check index)
drop index if exists instance_application_group_id_idx;

-- 4. Drop redundant single-column index on instance_status_history(group_id) (covered by composite index)
drop index if exists instance_status_history_group_id_idx;

-- 5. Drop redundant single-column index on instance_status_history(instance_id) (covered by composite index)
drop index if exists instance_status_history_instance_id_idx;

-- 6. Add index on instance(oem) to accelerate distribution reporting queries
create index if not exists instance_oem_idx on instance (oem);

-- +migrate Down

drop index if exists instance_oem_idx;

create index if not exists instance_status_history_instance_id_idx on instance_status_history (instance_id);
create index if not exists instance_status_history_group_id_idx on instance_status_history (group_id);
create index if not exists instance_application_group_id_idx on instance_application (group_id);
create index if not exists instance_application_instance_id_idx on instance_application (instance_id);
create index if not exists instance_application_instance_id_application_id_idx on instance_application (instance_id, application_id);
