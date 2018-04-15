
-- +migrate Up
create table passwords
(
	id int unsigned auto_increment
		primary key,
	user_id int unsigned not null,
	password varchar(32) null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	code varchar(100) null comment 'Код подтверждения'
);

create index passwords_code_index
	on passwords (code)
;

create index passwords_user_id_index
	on passwords (user_id)
;

create index passwords_password_index
	on passwords (password)
;



-- +migrate Down
drop table passwords;