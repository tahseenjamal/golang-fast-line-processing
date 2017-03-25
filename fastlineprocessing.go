// +build

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func line_processing(line string) {

	//print or do any processing of the line that has been read
	fmt.Println(line)

}

func main() {

	if len(os.Args) < 3 {

		fmt.Printf("\nError: Please pass two parameters: filename and buffer size\n\n")
		os.Exit(0)
	}

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
		buffer_size, _ := strconv.Atoi(os.Args[2])

		//initialize buffer data and reuse it by trimming method
		buffer_data := make([]byte, buffer_size)

		for {

			//get received data size
			size_received, _ := filehandle.Read(buffer_data)

			//trim the buffer to data size
			buffer_data = buffer_data[0:size_received]

			//if received size of data is not 0
			if size_received != 0 {

				//add to the buffer previous partial lines, which would have now their remaining part in this array
				buffer_data = []byte(partial_line + string(buffer_data))

				//count the number of sentence using newline
				array_count := len(strings.Split(string(buffer_data), "\n"))

				//if the last character in the byte array is not newline, it means last line is partial
				if buffer_data[size_received-1] != '\n' {

					//get the lines as array and note that the last line is partial
					lines := strings.Split(string(buffer_data), "\n")

					//depending on the buffer size initialised, we can have buffer holding the very first line also partially or fully
					//if at least the buffer size was big enough to hold the first line then
					//process all lines except for the last line

					if array_count > 0 {

						//get the partial line
						partial_line = lines[array_count-1]

						//trim out the partial line so you are left with only complete lines
						lines = lines[0 : array_count-1]

						for _, v := range lines {

							//process the complete lines
							line_processing(v)

						}

					} else {

						//if the buffer size chosen is too small to even take up first line then, it all starts with partial line only
						//and keep adding to partial line until the partial line becomes full line
						partial_line = string(buffer_data)

					}

				} else {

					//so in case the last character is newline, it means we have all complete lines in the array
					//split array basis newline
					lines := strings.Split(string(buffer_data), "\n")

					//now as last character is newline split basis newline would result in 1 extra array
					//this extra array has to be trimmed out
					lines = lines[0 : len(lines)-1]

					for _, v := range lines {

						line_processing(v)

					}

				}

				//received data size is 0 so break out, either file ended or start with nothing in file as the file was empty
			} else {

				break

			}

		}

	}

}
