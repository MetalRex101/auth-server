
-- +migrate Up
create table oauth_sessions
(
	id int unsigned auto_increment primary key,
	client_id int unsigned not null comment 'Клиент',
	user_id int unsigned null comment 'Пользователь',
	access_granted_at timestamp default CURRENT_TIMESTAMP not null comment 'Доступ открыт',
	access_expires_at datetime not null comment 'Доступ истекает',
	offset int(10) default '0' not null comment 'Разница во времени (сек, клиент - сервер)',
	code varchar(255) null comment 'Временный код',
	access_token varchar(100) null comment 'Код доступа',
	user_agent text null comment 'User Agent',
	remote_addr varchar(25) null comment 'REMOTE_ADDR',
	http_referer varchar(255) null comment 'HTTP_REFERER',
	constraint `UNIQUE`
		unique (client_id, user_id)
)
;

create index oauth_sessions_access_token_index
	on oauth_sessions (access_token)
;

create index oauth_sessions_access_expires_at_index
	on oauth_sessions (access_expires_at)
;

create index oauth_sessions_client_id_index
	on oauth_sessions (client_id)
;

create index oauth_sessions_user_id_index
	on oauth_sessions (user_id)
;

-- +migrate Down
drop table oauth_sessions;
