package main

import (
	"encoding/json"
	"fmt"
	"github.com/joliv/spark"
	"log"
	"os"
	"reflect"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)

	var in interface{}
	err := decoder.Decode(&in)
	if err != nil {
		log.Fatalln("Error parsing JSON input.", err)
	}

	switch in.(type) {
	case []interface{}:
		values := gatherNumericValues(in.([]interface{}))
		out := make(map[string]interface{})
		for key, array := range values {
			out[key] = spark.Line(array)
		}

		bytes, err := json.MarshalIndent(&out, "", "  ")
		if err != nil {
			log.Fatalln("Error serializing JSON output.", err)
		}
		fmt.Println(string(bytes))
	default:
		log.Fatalln("Unsupported JSON type:", reflect.TypeOf(in))
	}
}

func gatherNumericValues(jsonArray []interface{}) (numericValues map[string][]float64) {
	numericValues = make(map[string][]float64)
	for _, value := range jsonArray {
		switch value.(type) {
		case map[string]interface{}:
			for key, value := range value.(map[string]interface{}) {
				switch value.(type) {
				case float64:
					if array, ok := numericValues[key]; ok {
						numericValues[key] = append(array, value.(float64))
					} else {
						numericValues[key] = []float64{value.(float64)}
					}
				}
			}
		}
	}
	return
}
