# MAP

### Declaration

```go
var person map[string]string = map[string]string{}
person["name"] = "Putera"
person["gender"] = "male"

person2 := map[string]string{
    "name":   "Putera",
    "gender": "male",
}
```

### method

```go
len(map)            // get map data length
map(key)            // get value by key
map[key] = value    // set value by key
delete(map, key)    // delete value by key
make(map[typeKey][typeValue]) // create new map
```

```go
person3 := make(map[string]string)
person3["name"] = "dede"
person3["gender"] = "female"

fmt.Println(person3) // map[gender:female name:dede]

person4 := make(map[string]interface{})
person4["name"] = "dede"
person4["gender"] = "female"
person4["age"] = 34

fmt.Println(person4) // map[age:34 gender:female name:dede]
```
