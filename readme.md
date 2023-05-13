# MessagingService

### Technology stack
- Golang 1.20
- go-kit 0.12.0
- MySQL 8.0.30

### Local deployment
- build containers:
```shell
docker-compose up -d
```
- apply migrations from `migrations` folder:
```shell
docker exec -ti gomysql mysql -u *user* -p*pass* *dbname* < create_table_users.sql
docker exec -ti gomysql mysql -u *user* -*pass* *dbname* < create_table_messages.sql
```