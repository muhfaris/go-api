## Basic Model API

I create three type :
- started (main.go.started)
- using sqlite (main.go.sqlite)
- using mysql (main.go)

if you want to run type `started` or `sqlite`, please rename first !
### Requirement
#### Started
-  Gin (github.com/gin-gonic/gin)

#### Sqlite
- fmt
- Gin (github.com/gin-gonic/gin)
- Gorm (github.com/jinzhu/gorm)
- Sqlite (github.com/jinzhu/gorm/dialects/sqlite)

#### MySQL
- fmt
- Gin (github.com/gin-gonic/gin)
- MySQL (github.com/go-sql-driver/mysql)
- Gorm (github.com/jinzhu/gorm)

### Pattern URI
This just for type `sqlite` and `mysql`

```
GET    /person/                  
GET    /person/:id               
POST   /person                   
PUT    /person/:id               
DELETE /person/:id 
```

### HOW
#### Create person
``curl -i -X POST http://localhost:8080/person -d '{"Firstname":"Ali":"Lastname":"Rahman"}'``

Then open via browser, you will see result :
```
[{"ID":1,"Firstname":"Ali","Lastname":"Rahman"}]
```

#### Update person
``curl -i -X PUT http://localhost:8080/person/1 -d '{"Firstname":"Rahman":"Lastname":"Ali"}'``

Then open via browser, you will see result :
```
[{"ID":1,"Firstname":"Rahman","Lastname":"Ali"}]
```

#### Delete person 
``curl -i -X DELETE http://localhost:8080/person/1``

Then open via browser, you will see result :
```
[]
```