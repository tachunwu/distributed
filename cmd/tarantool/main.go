package main

import (
	"fmt"
	"log"

	"github.com/tarantool/go-tarantool"
)

func main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	resp, err := conn.Insert(
		"tester",
		[]interface{}{4, "ABBA", 1972},
	)
	fmt.Println(resp.Code)
	fmt.Println(resp.Data)

	resp, err = conn.Select(
		"tester", "primary", 0, 1, tarantool.IterEq,
		[]interface{}{4},
	)
	fmt.Println(resp.Code)
	fmt.Println(resp.Data)

	resp, err = conn.Select(
		"tester", "secondary", 0, 1, tarantool.IterEq,
		[]interface{}{"ABBA"},
	)
	fmt.Println(resp.Code)
	fmt.Println(resp.Data)
}
