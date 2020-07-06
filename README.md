# multilogger
Multilogger practice problem.Contains std,sys,database,file logger.
## Cloning
```
git clone https://github.com/aTTiny73/multilogger.git
```
## Database setup

To setup database first you need to instal mysql-server.
To get the exact same table as me, inside the mysql shell, type these commands :
```
CREATE DATABASE LOGGER;
USE LOGGER;
CREATE TABLE Log
(
      id int NOT NULL AUTO_INCREMENT,
    PREFIX varchar(255) NOT NULL,
    DATE varchar(255) NOT NULL,
    TIME varchar(255) NOT NULL,
    TEXT varchar(255) NOT NULL,
    PRIMARY KEY (id)
);
```
Now its time to setup user you can do that by running this command in mysql shell:

```
CREATE USER 'Testuser'@'localhost' IDENTIFIED BY '12345678';
```
Now we need to grant all privileges to user so he can add to the tabel delete etc.
```
GRANT ALL PRIVILEGES ON LOGGER.Log TO 'Testuser'@'localhost';
```
