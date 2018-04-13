# Freelance

## Установка MySQL сервера
После установки создаем базу freelance и пользователя:
- $mysql -u root -p
- mysql> CREATE DATABASE freelance;
- mysql> GRANT ALL PRIVILEGES ON freelance.* TO dbuser_f@localhost IDENTIFIED BY 'dbpass_f';
- mysql> exit;

## Установка PostgreSQL сервера
Искать в Google

## Установка GRPC

## Настройка migrations of DB
Миграции БД - это типа контроль версии схемы базы данных. Т.е. файлы или команды, описание создания БД, изменения структуры (добавление/удаление/изменение колонок, индексов ...), потом можно поделится этими изменениями. Получив эти файлы(команды) модно воспроизвести у себя в БД и получишь последние изменения(обновления) БД. 
- для начала тянем библиотеку: go get github.com/mattes/migrate
- создадим две папки внутри migrations, для двух типов БД: mysql и postresql
- сейчас миграция работает для MySQL, запускается в main.go. Потом изменим.

## Структура папок
Структура папок (будем дополнять по мере появления файлов):
|__api
|__db
|__freelance
|__model
|__server
|__migration
    |_mysql
       |__001_users.up.sql
       |__001_users.down.sql
       |__002_tasks.up.sql
       |__002_tasks.down.sql
       |__003_billings.up.sql
       |__003_billings.down.sql
    |_pgsql
       |__...
|__cmd
    |__freelance
        |__main.go
|__Makefile
|__README.md