## Basic API

### Requirement
- encoding/json
- fmt
- log
- net/http
- time
- mux (github.com/gorilla/mux)


### URI Pattern
| HTTP        | Pattern          | 
| ------------- |:-------------:| 
| GET     | /todos | 
| GET      | /todos/{todoId}      | 


### Example Output

```
[{
	"Name":"Write presentation",
	"Completed":false,
	"Due":"0001-01-01T00:00:00Z"
},{
	"Name":"host meetup",
	"Completed":false,
	"Due":"0001-01-01T00:00:00Z"
}]
```