package fundamentals

import "fmt"

type A struct {
	s string
	i int
	f float32
}

// Test out formating of printf method
func Print() {
	fmt.Println("Number")
	fmt.Printf("%v", 1)                                // 1
	fmt.Printf("%#v", 1)                               // 1
	fmt.Printf("%T", 1)                                // int
	fmt.Printf("a %% sign and type %T value %v", 1, 1) // a % sign and type int value 1
	fmt.Printf("%b", 10)                               // 1010
	fmt.Printf("%c", 78)                               // N
	fmt.Printf("%d", 10)                               // 10
	fmt.Printf("%o", 10)                               // 12
	fmt.Printf("%x", 30)                               // 1e
	fmt.Printf("%X", 30)                               // 1E
	fmt.Print(fmt.Sprintf("%x", 30))                   // 1e
	fmt.Printf("%U", 30)                               // U+001E
	fmt.Printf("%.2f", 3.48231)                        // 3.48
	fmt.Printf("%6.2f", 3.48231)                       //   3.48

	fmt.Println("Struct")
	s := A{"string value", 10, 32.42812}
	fmt.Printf("%v", s)  // {string value 10 32.42812}
	fmt.Printf("%#v", s) // fundamentals.A{s:"string value", i:10, f:32.42812}
	fmt.Println(s)       // {string value 10 32.42812}

	fmt.Println("Slice")
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl1 := sl[2:5]
	sl2 := make([]int, 3, 5)
	fmt.Print(sl)                                       // [1 2 3 4 5 6 7 8 9]
	fmt.Printf("%v", sl)                                // [1 2 3 4 5 6 7 8 9]
	fmt.Printf("%T", sl)                                // []int
	fmt.Print(sl1)                                      // [3 4 5]
	fmt.Printf("%T", sl1)                               // []int
	fmt.Print(sl2)                                      // [0 0 0]
	fmt.Printf("length %d, cap %d", len(sl2), cap(sl2)) // length 3, cap 5
	fmt.Printf("%p", sl2)                               // 0xc00001c120

	fmt.Println("Map")
	m := map[string]int{
		"value1": 1,
		"value2": 2,
		"value3": 3,
	}
	fmt.Print(m)                         // map[value1:1 value2:2 value3:3]
	fmt.Printf("%#v", m)                 // map[string]int{"value1":1, "value2":2, "value3":3}
	fmt.Println("Value of the map: ", m) // Value of the map:  map[value1:1 value2:2 value3:3]
}
