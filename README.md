# Freelance

## Установка MySQL сервера

После установки создаем базу freelance и пользователя:

- sudo apt-get install mysql-server mysql-client
- $mysql -u root -p
- mysql> CREATE DATABASE freelance;
- mysql> GRANT ALL PRIVILEGES ON freelance.* TO dbuser_f@localhost IDENTIFIED BY 'dbpass_f';
- mysql> exit;

## Установка PostgreSQL сервера

Искать в Google

- устанавливаем PostgreSQL

sudo apt-get install postgresql postgresql-contrib

- запустим оболочку

sudo -u postgres psql

- создаем пользователя

create user dbuser_f with password 'dbpass_f';

- создаем базу данных

create database freelance;

- выдаем права доступа к базе

grant all privileges on database freelance to dbuser_f;

- выходим из оболочки

\q

## Установка GRPC

## Настройка migrations of DB

Миграции БД - это типа контроль версии схемы базы данных. Т.е. файлы или команды, описание создания БД, изменения структуры (добавление/удаление/изменение колонок, индексов ...), потом можно поделится этими изменениями. Получив эти файлы(команды) модно воспроизвести у себя в БД и получишь последние изменения(обновления) БД.

- для начала тянем библиотеку: go get github.com/mattes/migrate
- создадим две папки внутри migrations, для двух типов БД: mysql и postresql
- сейчас миграция работает для MySQL и для PgSQL, запускается в main.go. Расскоментировать нужную строку, вторую закоментить. Перед запуском main.go посмотрите Makefile , там есть команда export DATABASE_URL ..., нужно скопипастить и запустить в терминале вручную , для своей БД.
- в результате выполнения, у вас появятся таблицы в базе. Если захотите откатить назад, тогда в main.go изменить MigrateUp() на MigrateDown(), в результате все таблицы удалятся из БД.