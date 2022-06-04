package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var sha512_flag = flag.Bool("sha512", false, "to get sha512 values")
var sha384_flag = flag.Bool("sha384", false, "to get sha384 values")

func main() {
	fmt.Printf("Enter the value to get ssh: ")
	ip_reader := bufio.NewReader(os.Stdin)
	ip_str, _ := ip_reader.ReadString('\n')
	ip_str = strings.TrimSpace(ip_str)
	flag.Parse()
	if *sha512_flag {
		sha512_val := sha512.Sum512([]byte(ip_str))
		fmt.Printf("sha512: %x\n", sha512_val)
	} else if *sha384_flag {
		sha384_val := sha512.Sum384([]byte(ip_str))
		fmt.Printf("sha384: %x\n", sha384_val)
	} else {
		sha256_val := sha256.Sum256([]byte(ip_str))
		fmt.Printf("sha256: %x\n", sha256_val)
	}
}
