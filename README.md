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

## Настройка migrations of DB

Миграции БД - это типа контроль версии схемы базы данных. Т.е. файлы или команды, описание создания БД, изменения структуры (добавление/удаление/изменение колонок, индексов ...), потом можно поделится этими изменениями. Получив эти файлы(команды) можно воспроизвести у себя в БД и получишь последние изменения(обновления) БД.

- для начала тянем библиотеку: go get github.com/mattes/migrate
- создадим две папки внутри migrations, для двух типов БД: mysql и postresql
- сейчас миграция работает для MySQL и для PgSQL, запускается в main.go. Расскоментировать нужную строку, вторую закоментить. Перед запуском main.go посмотрите Makefile , там есть команда export DATABASE_URL ..., нужно скопипастить и запустить в терминале вручную , для своей БД.
- в результате выполнения, у вас появятся таблицы в базе. Если захотите откатить назад, тогда в main.go изменить MigrateUp() на MigrateDown(), в результате все таблицы удалятся из БД.

## Транзакции БД

Транзакция — это операция, состоящая из одного или нескольких запросов к базе данных. Суть транзакций — обеспечить корректное выполнение всех запросов в рамках одной транзакции, а так-же обеспечить механизм изоляции транзакций друг от друга для решения проблемы совместного доступа к данным.

Любая транзакция либо выполняется полностью, либо не выполняется вообще.

Пример:
START TRANSACTION;
UPDATE user_account SET allsum=allsum + 1000 WHERE id='1';
UPDATE user_account SET allsum=allsum - 1000 WHERE id='2';
COMMIT;

В нашем коде используя библиотеку sqlx:
tx := m.conn.MustBegin()
tx.MustExec(
    "UPDATE `billings` SET sender=?, reciever=?, amount=?, time_bill=?, task_id=?, btype=? WHERE id=?",
    s.Sender, s.Reciever, s.Amount, s.TimeBill, s.TaskID, s.BillingType, s.ID,
)
err := tx.Commit()

## Установка GRPC

- вставим все как описано в офф.сайте
- api.proto это описание protobuf пока создал 4 метода для работы с таблицей user, необходимо будет по аналогии создать для других таблиц
- запускаем в терминале make api , в результате в папке api сгенерится два файла: api.pb.go и api.pb.gw.go
- создаем вручную там же handler.go, наши функции обработчики, пока для вышеобозначеных 4 методов, нужно будет дополнить
- в папке server создаем наш сервер rest api

## Пример запуска

- необходимо сначала экспортировать переменные окружения

 export DB_TYPE="mysql" (либо postgres)
 export DO_MIGRATION="No" (либо Yes)
 export DATABASE_URL=....

- пока работает только методы для User и Billing
- go run main.go
- и в терминале вводим команду и видим возвращаемую id юзера:

curl -X POST -k http://localhost:7778/api/user -d '{"name":"firstuser", "utype":"client", "balance":1}'

- в базе создается наш юзер
- в терминале вводим команду, для удаления пользователя по номером 4

curl -X DELETE -k http://localhost:7778/api/user/4

- в терминале вводим команду, для обновления пользователя под номером 3

curl -X POST -k http://localhost:7778/api/user/3 -d '{"name":"tom soyer", "utype":"client", "balance":15}'

- аналогично выше указанным командам можно проделать для billing
- какие методы есть для api смотреть в api.proto
