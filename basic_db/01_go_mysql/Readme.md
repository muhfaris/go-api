## How to connect to MySQL
### Prepare
- Create database
- Create table `user` with schema like :
```
+-------+-------------+------+-----+---------+----------------+
| Field | Type        | Null | Key | Default | Extra          |
+-------+-------------+------+-----+---------+----------------+
| id    | int(11)     | NO   | PRI | NULL    | auto_increment |
| name  | varchar(25) | YES  |     | NULL    |                |
| email | varchar(25) | YES  |     | NULL    |                |
+-------+-------------+------+-----+---------+----------------+
```

- Insert data into table

### Configuration Database
Location of configuration database in `database/database.go`

### Main
In function `main` just read data from database `db1` table `user`. Data from database is save into variabel data (slice type of struct User)