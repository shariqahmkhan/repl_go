package main
/*
import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

  // read whole content and not EOF
	content, err := ioutil.ReadFile("abcd.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}
*/

import (
"os"
"bufio"
"fmt"
"log"

)
// read in bytes 
func main(){

// first open a file
file, err := os.Open("abcd.txt") // For read access.
if err != nil {
	log.Fatal(err)
}

// each package as an interface implementation (such as Scanner, Reader)
// Generally, you perform Reader, NewScanner(Scanner) then either prints or scans it line by line

// since os.Open returns Reader, so bufio.Reader not required
scanner := bufio.NewScanner(file)

// again in type Scanner, we have Split to perform splitting operation
scanner.Split(bufio.ScanLines)

// then we have Scan to Scan next each either by line or bytes
// then infinite looping until EOF
// EOF is befcause of Scan()
for scanner.Scan(){
  fmt.Println(scanner.Text())
}

/*
Referece:
https://www.socketloop.com/references/golang-bufio-scanlines-function-example
https://golang.org/pkg/bufio/#ScanBytes
*/
}
