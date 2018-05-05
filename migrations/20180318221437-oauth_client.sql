
-- +migrate Up
create table oauth_clients
(
	id int unsigned auto_increment
		primary key,
	status tinyint(1) unsigned default '0' not null comment 'Вкл./выкл.',
	name varchar(200) not null comment 'Название',
	client_id varchar(200) not null comment 'client_id',
	client_secret varchar(200) not null comment 'client_secret',
	template varchar(50) null comment 'Внешний вид',
	ip varchar(96) not null comment 'IP-адрес клиента',
	url varchar(255) not null comment 'Адрес сайта',
	scope varchar(255) null comment 'Права доступа'
);

create index oauth_client_client_id_index
	on oauth_clients (client_id)
;

create index oauth_client_client_secret_index
	on oauth_clients (client_secret)
;

create index oauth_client_status_index
	on oauth_clients (status)
;

-- +migrate Down
drop table oauth_clients;
