package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// could look more from shell:
	// curl -i https://api.github.com/users/tebeka
	// or save it off to file from curl
	url := "https://api.github.com/users/tebeka"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", resp.Status)
	}
	fmt.Println("**** All is well")
	fmt.Println("content-type:", resp.Header.Get("Content-Type"))
	fmt.Println("---------------------")
	if written, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: %s", err)
	} else {
		fmt.Println("\nWrote: ", written, " bytes")
	}
}

type reply struct {
	Name         string
	Public_Repos int
}

/* REST API:
CRUD: Create, Retrievex, Update, Delete
Create: POST
Retrieve: GET
Update: PUT
Delete: DELETE

JSON <-> Go
string : string
null : nil
	but in Go a string can't be nil...some ways around but in general that is the case
boolean : bool
number : float64 (default), float32, int8, int16, int32, int64, uint8, ...
array : []T, []any
object : struct, map[string]any

encoding/json
JSON -> io.Reader -> Go
	use Decoder
Go -> io.Writer -> JSON
	use Encoder
JSON -> []byte -> Go
	use Unmarshal
Go -> []byte => JSON
	use Marshal

*/
