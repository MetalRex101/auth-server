
-- +migrate Up
create table emails
(
	id int unsigned auto_increment primary key,
	user_id int unsigned not null,
	oauth varchar(255) null comment 'К каким внешним oauth-сервисам относится',
	email varchar(100) not null comment 'Адрес e-mail',
	status tinyint(1) unsigned default 1 not null comment 'Вкл./выкл.',
	is_default tinyint(1) unsigned default 0 not null comment 'Основной',
	confirm_date datetime null comment 'Подтвержден',
	code varchar(90) null comment 'Код активации',
	created_at timestamp default CURRENT_TIMESTAMP not null comment 'Создан'
);

create index emails_confirm_date_index
	on emails (confirm_date)
;

create index emails_email_index
	on emails (email)
;

create index emails_oauth_index
	on emails (oauth)
;

create index emails_user_id_index
	on emails (user_id)
;

create index emails_status_index
	on emails (status)
;

-- +migrate Down
drop table emails;
