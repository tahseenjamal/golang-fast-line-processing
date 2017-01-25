// +build

package main

import (
	/*"bytes"*/
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	filehandle, err := os.Open(os.Args[1])

	defer filehandle.Close()

	if err == nil {

		//initialise partial line variable
		partial_line := ""

		buffer_size := 1024
		buffer_data := make([]byte, buffer_size)

		for {

			/*buffer_data := make([]byte, 5)*/

			size_received, err := filehandle.Read(buffer_data)
			buffer_data = buffer_data[0:size_received]

			if err == io.EOF || size_received == 0 {

				if partial_line != "" {
					fmt.Println(partial_line)
				}
				break
			}

			//Add any previous partial byte to this []byte so that when we split, the inital part of has previous last partial data
			buffer_data = []byte(partial_line + string(buffer_data))

			//split the []byte basis \n

			buffer_line_array := strings.Split(string(buffer_data), "\n")
			/*buffer_line_array[0] = partial + b2[0]*/

			//flush partial array
			partial_line = ""

			//if array string is having len not matching newline means we have partial line in the last cell
			if len(buffer_line_array) != strings.Count(string(buffer_data), "\n") {

				//Check if we have only 1 cell or more than that
				if len(buffer_line_array) > 1 {

					//if more than 1 cell, then get the partial line in partial variable so that it can be added in the beginning to the next buffer read
					partial_line = buffer_line_array[len(buffer_line_array)-1]

					//print all except for the partial line
					for index := 0; index <= len(buffer_line_array)-2; index++ {
						fmt.Println(buffer_line_array[index])
					}

				} else {

					//if only 1 line then we store that partial line in partial variable, ready to be added to the next buffer read

					partial_line = buffer_line_array[0]
				}

			} else {

				//this condition when cells match newline, means no partial data

				partial_line = ""

				for index := 0; index < len(buffer_line_array)-1; index++ {
					fmt.Println(buffer_line_array[index])
				}

			}

		}

	}

}
