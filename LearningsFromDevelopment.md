# Learnings from Development in GO

```go
// CASE 1 : 
type IdName struct {
   Id   *int    `json:"id,omitempty"` // if id,name = nil it will skip it from json
   Name *string `json:"name,omitempty"`
}


// CASE 2 : 
type IdName struct {
   Id   *int    `json:"id,omitempty"` // if id,name = nil it will skip it from json
   Name *string `json:"name,omitempty"`
}

type GetDealGeneralData struct {
   Licensee IdName `json:"licensee" validate:"omitempty"`
}
```

if for case 2 the json is like 

```json
"licensee": {
   "id": null,
   "name": null
}
```

then for case 1 it is like : 

```json
"licensee": {}
// i.e., json:"id,omitempty" will skip the entry from json generated if that field is null
```