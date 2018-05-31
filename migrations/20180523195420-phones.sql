
-- +migrate Up
create table phones
(
	id int unsigned auto_increment
		primary key,
	user_id int unsigned not null,
	oauth varchar(255) null comment 'К каким внешним oauth-сервисам относится',
	phone bigint unsigned not null comment 'Номер телефона',
	status tinyint(1) default '1' not null comment 'Вкл./выкл.',
	is_default tinyint(1) default '0' not null comment 'Основной',
	confirm_date datetime null comment 'Подтвержден',
	code varchar(90) null comment 'Код активации',
	created_at datetime default CURRENT_TIMESTAMP not null comment 'Время создания'
)
;

create index phones_confirm_date_index
	on phones (confirm_date)
;

create index phones_oauth_index
	on phones (oauth)
;

create index phones_phone_index
	on phones (phone)
;

create index phones_status_index
	on phones (status)
;

create index phones_user_id_index
	on phones (user_id)
;

-- +migrate Down
drop table phones;