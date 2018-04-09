Freelance

Установка MySQL сервера

Установка PostgreSQL сервера

Установка GRPC

Настройка migrations of DB
- для начала тянем библиотеку: go get github.com/mattes/migrate
- создадим две папки внутри migrations


СТруктура папок (будем дополнять по мере появления файлов):
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
       |__001_users.up.sql
       |__001_users.down.sql
       |__002_tasks.up.sql
       |__002_tasks.down.sql
       |__003_billings.up.sql
       |__003_billings.down.sql
|__cmd
    |__freelance
        |__main.go
|__Makefile
|__README.md