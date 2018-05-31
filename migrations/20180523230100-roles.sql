
-- +migrate Up
create table roles
(
	id int unsigned auto_increment
		primary key,
	code varchar(50) not null comment 'Код роли',
	name varchar(50) not null comment 'Название',
	description varchar(255) null comment 'Описание',
	synthetic tinyint(1) default '0' not null comment 'Синтетическая роль',
	created_at timestamp default CURRENT_TIMESTAMP null comment 'Время создания',
	updated_at datetime null comment 'Время обновления',
	creator_id int unsigned null comment 'Создатель записи',
	editor_id int unsigned null comment 'Редактор записи',
	status tinyint(1) unsigned default '0' null comment 'Вкл./выкл.',
	constraint code
		unique (code),
	constraint name
		unique (name)
)
comment 'Роли пользователей'
;

create index roles_creator_id_index
	on roles (creator_id)
;

create index roles_editor_id_index
	on roles (editor_id)
;

-- +migrate Down
drop table roles;
