package main

import (
	"encoding/json"
	"fmt"
)

// since c is small-caps then it will be available in this file only
type course struct {
	Name     string   `json:"courseName"`
	Price    int      // since no json tag exact struct name is used as key in json
	Platform string   `json:"website"`
	password string   `json:"-"`              // - in json tag removes the key from struct json map
	Tags     []string `json:"tags,omitempty"` // if it remains nil then it is omitted from json
}

func main() {
	fmt.Println("json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"Go lang Bootcamp", 299, "google.com", "abc123", []string{"go", "programming"}},
		{"Java Bootcamp", 199, "oracle.com", "pswd123", nil},
		{"Python Bootcamp", 499, "python.org", "cryptoHEX", []string{"python", "not snake", "programming"}},
	}

	// packaging struct data as json
	jsonParsed1, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json : \n%s\n", jsonParsed1)

	// packaging struct data as json
	jsonParsed2, err := json.MarshalIndent(lcoCourses, "", "\t") // \t is tab in indent
	if err != nil {
		panic(err)
	}
	fmt.Printf("json : \n%s\n", jsonParsed2)
}

func DecodeJson() {

	// any data from web is received as a byte array
	jsonWebData := []byte(`
		{
			"courseName":"Go lang Bootcamp",
			"Price":299,
			"website":"google.com",
			"tags":["go","programming"]
		}
	`)
	jsonDecoded := course{}

	isValid := json.Valid(jsonWebData)
	if isValid {
		fmt.Println("json is valid")

		// passing ref(address) since json.Unmarshal doesnt return parsed data but it modifies the var
		// provided json.Unmarshal needs to modify the actual variable to populate it with the parsed
		// data. On passing the variable by value, json.Unmarshal would only modify a copy of the
		// variable, and the original variable would remain unchanged.
		err := json.Unmarshal(jsonWebData, &jsonDecoded)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", jsonDecoded)
		fmt.Printf("%+v\n", jsonDecoded)
	} else {
		fmt.Println("INVALID JSON")
	}

	// some cases where we just want to add data to a key value pair.
	var keyValJsonParsed map[string]interface{}
	// in json 1st level key is string but value can have variable types (int, string, bool, object, array)
	err := json.Unmarshal(jsonWebData, &keyValJsonParsed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", keyValJsonParsed)

	for k, v := range keyValJsonParsed {
		fmt.Printf("key: %v | val: %v | type: %T\n", k, v, v)
	}
	// note that for Price in struct it is int but on unmarshalling it in interface it will converted to float

}
