build:
	docker build -t test/todo_build .
	docker run -e MYSQL_HOST=192.168.0.104 -e MYSQL_USER=root -e MYSQL_PASSWORD=faris -e MYSQL_DBNAME=skyshi -p 8090:3030 test/todo_build
