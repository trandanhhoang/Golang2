UPDATE mysql.user SET Password=PASSWORD('bongbong123') WHERE User='root';

ALTER USER 'root'@'localhost' IDENTIFIED BY 'bongbong123';
GRANT ALL PRIVILEGES ON _ . _ TO 'root'@'localhost';

docker run --name=mysql2 --restart on-failure -e MYSQL_ROOT_HOST='172.17.0.1' -d -p 3306:3306 mysql mysql-server:8.0

docker run --name some-mysql -e MYSQL_ROOT_HOST=% -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest

docker logs mysql4 2>&1 | grep GENERATED
docker exec -it mysql4 mysql -uroot -p

CREATE USER 'username'@'172.17.0.1' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON _._ TO 'username'@'172.17.0.1' WITH GRANT OPTION;
