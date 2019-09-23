package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	//tls.Config{}
	c, err := ftp.Dial("114.116.1.186:21", ftp.DialWithTimeout(20*time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("dial success")

	err = c.Login("ftptest", "12345678")
	if err != nil {
		log.Fatal("login:", err)
		return
	}

	file, err := os.Open("/tmp/app.log")
	if err != nil {
		log.Fatal(err)
		return
	}

	/*if err := c.MakeDir("/tmp/data"); err != nil {
		log.Fatal("make dir error: ", err)
		return
	}*/
	err = c.Stor("./test.txt", file)
	if err != nil {
		log.Fatal("stor error:", err)
		return
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
