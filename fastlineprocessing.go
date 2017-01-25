// +build

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//reading file name from the command line argument
	filename := os.Args[1]

	//creating file handle
	filehandle, err := os.Open(filename)

	//close file handle before exiting main()
	defer filehandle.Close()

	//enter for processing only if filehandle creation for read successful
	if err == nil {

		//initialise partial line variable
		partial_line := ""

		//buffer size of buffer data
		buffer_size := 1024

		//initialize buffer data and reuse it by trimming method
		buffer_data := make([]byte, buffer_size)

		for {

			//read buffer data
			size_received, err := filehandle.Read(buffer_data)

			//trim buffer to the size of data received
			buffer_data = buffer_data[0:size_received]

			//Plan exit if EOF or size received 0
			//I think I should make err != nill as already am checing size. err should be use to catch for any other possible errors
			if err == io.EOF || size_received == 0 {

				//If partial line from previous read is not null, print it too before exiting
				if partial_line != "" {
					fmt.Println(partial_line)
				}

				//exit loop and this results in exit of program as well
				break
			}

			//Add any previous partial byte to this []byte so that when we split, the inital part of has previous last partial data
			buffer_data = []byte(partial_line + string(buffer_data))

			//split the []byte basis \n

			buffer_line_array := strings.Split(string(buffer_data), "\n")
			/*buffer_line_array[0] = partial + b2[0]*/

			//flush partial first
			partial_line = ""

			//if array string is having length that does not match newline count, then it means we have partial line in the last cell
			if len(buffer_line_array) != strings.Count(string(buffer_data), "\n") {

				//Check if we have only 1 cell or more than that
				if len(buffer_line_array) > 1 {

					//if more than 1 cell, then get the partial line in partial variable so that it can be added in the beginning to the next buffer read
					totalsize := len(buffer_line_array)

					//partial line is in the last cell. Note cells start from 0 to totalsize-1, where totalsize is the number of cells
					partial_line = buffer_line_array[totalsize-1]

					//print all except for the partial line
					for index := 0; index <= len(buffer_line_array)-2; index++ {

						fmt.Println(buffer_line_array[index])
					}

					//enter if buffer read is of 1 cell only, which means a file with single line
				} else {

					//if only 1 line then we store that partial line in partial variable, ready to be added to the next buffer read
					partial_line = buffer_line_array[0]
				}

				//this condition when cells match newline, means no partial data
			} else {

				for index := 0; index < len(buffer_line_array)-1; index++ {

					//print the complete lines aggregated from the current read buffer data.
					fmt.Println(buffer_line_array[index])
				}

			}

		}

	}

}
