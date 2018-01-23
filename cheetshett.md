## Cheetshet GOLANG
### Variabel
Declaration to zero value ``var NAME TYPE``
Declaration and assigning a value ``NAME := VALUE``
Assigning to a previously declared variabel  ``Name=VALUE``

Default value variabel if not declaration
- integer   = 0
- booleans  = false
- strings   = ""

### Handy short variabel 
use ``:=``, like ``power:=9000``

### Pointer
Can use ``New()`` or ``&``  or ``*``
ex :
``
goku := new(Saiyan)

vs

goku := &Saiyan{}

---------------------
goku := new(Saiyan)
goku.name = "Goku"
goku.power = 9001

vs

goku := &Saiyan(
    name: "Goku",
    power: 9001,
)
``
### Array
How to declare in array
`
var a [2] string
`
`
a := [2]string{"hello","world"}
`
``
a := [...]string{"hello","world"}
``

### Slice
``
p := []int{1,2,3,4,5,6}
``
``
a := make([]string,3)
``
Append
``
a := make([]string{})
a = append(a, "Indonesia")
``


### Maps
Maps in Go are what other languages call hashtables or directionaries

ex:
```
func main() {
    lookup := make(map[string]int)
    lookup["goku"] = 9001
    power, exists := lookup["vegeta"]

    // prints 0, false
    // 0 is the default value for an integer
    fmt.Println(power, exists)
}
```
Map as field of a structure
```
type Saiyan struct {
    Name string
    Friends map[string]*Saiyan
}
```

### Package
- Package names follow the directory structure of your GO workspace

### Access Offline Documentation
- $ godoc -http=:6060
_* bind to port 6060_

### Note
- Go isn't object-oriented (OO) language, dont worry you will be notice a lot of
  similarities

