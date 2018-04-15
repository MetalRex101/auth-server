
-- +migrate Up
create table users
(
	id int unsigned auto_increment primary key,
	created_at timestamp default CURRENT_TIMESTAMP null comment 'Время создания',
	updated_at datetime null comment 'Время обновления',
	changed_at datetime null comment 'Время изменения',
	last_visit datetime null comment 'Последний вход',
	creator_id int unsigned null comment 'Создатель записи',
	editor_id int unsigned null comment 'Редактор записи',
	status tinyint(1) unsigned default '0' null comment 'Вкл./выкл.',
	merged text null comment 'Объединенные пользователи',
	nickname varchar(50) null comment 'Отображаемое имя',
	first_name varchar(100) null comment 'Имя',
	father_name varchar(100) null comment 'Отчество',
	last_name varchar(100) null comment 'Фамилия',
	birth_date date null comment 'Дата рождения',
	gender enum('m', 'f') null comment 'Пол',
	language varchar(255) null
);

create index users_creator_id_index
	on users (creator_id)
;

create index users_editor_id_index
	on users (editor_id)
;

-- +migrate Down
drop table users;