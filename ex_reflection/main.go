package main

import "fmt"
import "reflect"

func main() {

	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	var smaple1 struct {
		Offset uint `long:"offset" description:"Offset" gotohell:"124"`
	}
	st1 := reflect.TypeOf(smaple1)
	f2 := st1.Field(0)
	fmt.Println(f2)
}

// s := smaple{}
//	st = reflect.TypeOf(sample1)

//}
