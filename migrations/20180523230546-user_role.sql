
-- +migrate Up
create table user_role
(
	id int unsigned auto_increment
		primary key,
	created_at timestamp default CURRENT_TIMESTAMP null,
	user_id int unsigned not null,
	role_id int unsigned not null,
	constraint user_role_user_id_role_id_unique
		unique (user_id, role_id)
)
;

create index user_role_role_id_index
	on user_role (role_id)
;

create index user_role_user_id_index
	on user_role (user_id)
;

-- +migrate Down
drop table user_role;