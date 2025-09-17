package servicefetch

import (
	"encoding/json"
	"fmt"
	"go-git/clone"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

var filetype = []string{".yaml", ".properties"}

func Retrieve_data(c *gin.Context) {

	filename := c.Param("filename")
	environment := c.Param("environment")

	var filenames = []string{filename, filename + "/" + filename}

	for _, files_folders := range filenames {

		for _, extension := range filetype {
			openfile, err := os.OpenFile(clone.Path+"/"+files_folders+"-"+environment+extension, os.O_RDONLY, 0600)
			if err != nil {
				continue
			}

			file_contents, err := io.ReadAll(openfile)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(file_contents))

			//get the file type, whether yaml or properties
			file_type := strings.Split(openfile.Name(), ".")

			if file_type[2] == "yaml" || file_type[2] == "yml" {

				var data interface{}

				//translates the data into a map
				err = yaml.Unmarshal(file_contents, &data) //file contents have to be of type []byte
				if err != nil {
					log.Fatalf("Error unmarshaling YAML: %v", err)
				}

				// 3. Marshal the Go data structure into JSON
				// Use json.MarshalIndent for pretty-printed JSON output
				jsonData, err := json.MarshalIndent(data, "", "  ")
				if err != nil {
					log.Fatalf("Error marshaling to JSON: %v", err)
				}

				// formatteddata, err := json.Marshal(data)
				// fmt.Println(string(formatteddata))

				c.Data(http.StatusOK, "application/json", jsonData)
			} else {
				if file_type[2] == "properties" {

				}
			}

			defer openfile.Close()
		}
	}

}
