package main 

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
)

 func reqConn(x int) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(x))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(body))
	res.Body.Close()
}

func main() {
	for i:= 1; i <= 100; i++ {
		go reqConn(i)
	}
	var input string
    fmt.Scanln(&input)
}