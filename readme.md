# MessagingService

### Technology stack
 - Golang 1.20
 - MySQL 8.0.30
 - go-kit 0.12.0

### Local deployment
- build containers:
```shell
docker-compose up -d
```
- apply migrations from `migrations` folder:
```shell
docker exec -ti gomysql mysql -u test -ptest messaging < create_table_users.sql
docker exec -ti gomysql mysql -u test -ptest messaging < create_table_messages.sql
```