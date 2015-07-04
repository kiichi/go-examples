package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//--------------------------------------------------------
// Start Server:
// e.g. ~/bin/redis-3.0.2/src/redis-server
//
// running redis client
// redis-cli
//--------------------------------------------------------
func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	conn.Do("SET", "name", "Kiichi")
	str, err := redis.String(conn.Do("GET", "name"))
	fmt.Println(str)
	defer conn.Close()
}
