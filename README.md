# udplistener
Simple UDP listen server in Go (with logging functionality)

## How to run

### Cloning
You can "git clone" my repo with :

```
git clone https://github.com/TRedzepagic/udplistener.git
```
Then run with :

```
go run main.go "Port" 
```
## Database configuration (Taken from compositelogger configuration)

To setup the database you need to install the mysql-server, which you can look up online.
Since this program uses the configuration found in compositelogger for logging messages, the database configuration steps will be repeated here for clarity.

**NOTE:** Database is named "LOGGER" on mysql server, table is named "LOGS".

To get the exact same table as me, inside the mysql shell, type these commands :
```
CREATE DATABASE LOGGER;
USE LOGGER;
CREATE TABLE LOGS
(
    id int NOT NULL AUTO_INCREMENT,
    PREFIX varchar(255) NOT NULL,
    DATE varchar(255) NOT NULL,
    TIME varchar(255) NOT NULL,
    TEXT varchar(255) NOT NULL,
    PRIMARY KEY (id)
);
```
While on the server, you can create a user with this command :

```
"CREATE USER 'compositelogger'@'localhost' IDENTIFIED BY 'Mystrongpassword1234$';"
```
Then you need to grant the user access to our logging table, or else we will get an error :

```
"GRANT ALL PRIVILEGES ON LOGGER.LOGS TO 'compositelogger'@'localhost';"
```
Here we granted all privileges on our "LOGS" table to our user named "compositelogger".
