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

### Use
#### Pattern URI
Method | URI | 
------------ | -------------
GET | /users | Get all users
GET | /users/id | Get user by id
POST | /users | Create new user
POST | /users/id | Delete user by id

*_Create new user add parameter `Name` and `Email`_

### Example
#### GET
`curl -v http://localhost:9393/users`
`curl -v http://localhost:9393/users/1`

#### POST
`curl -i -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:9393/users -d "Name=Faris&Email=faris@email.com"`

#### DEL
`curl -i -H "Content-Type: application/x-www-form-urlencoded" -X DEL http://localhost:9393/users/1`

