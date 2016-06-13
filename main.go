package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./Dockerfile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "FROM alpine:edge")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "MAINTAINER Chuanjian Wang <me@ckeyer.com>")
	fmt.Fprintln(f)
	for I, J := 'a', 0; I <= 'z'; I++ {
		for i, j := 'a', 'A'-'a'; i <= 'z'; i++ {
			if J++; J > 125 {
				return
			}
			fmt.Fprintf(f, "RUN mkdir -p /opt/%s%s\n", string(I), string(i))
			fmt.Fprintf(f, "ENV %s%s=%s%s\n", string(I+j), string(i+j), string(I), string(i))
		}
	}
}
