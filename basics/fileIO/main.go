package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println


func main() {
  // basic file writing
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2,3,5,7,11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString((num + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
  
  
  /*
	Exacly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

	O_RDONLY : open the file read-only
	O_WRONLY : open the file write only
	O_RDWR   : open the file read-write
	
	The can be or'ed

    O_APEEND : append data to the file when writing
	O_CREATE : create a new file if none existis
	O_EXCL   : used with O_CREATE, file must not exist
	O_SYNC   : open for synchronous I/O
	O_TRUNC  : truncate regular writable file when openned
	*/
  
  _, err := os.Stat("data.txt") // verify is the data exists
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exists")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
