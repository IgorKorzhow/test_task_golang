# test_task_golang

1. in directory configs create file app.env on base app.env.example
2. Run command `docker compose up -d`
3. Open container api using `command docker-compose exec -it api bash`
4. run command in project directory `migrate -path migrations -database "mysql://user:ivalin60@tcp(mysql:3306)/test_task_golang" up`
5. open page `localhost:8080/ping` to check that service is running