package microservicefetch

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var filetype = []string{".yaml", ".properties"}

func Retrieve_data(c *gin.Context) {

	filename := c.Param("filename")
	environment := c.Param("environment")

	for _, extension := range filetype {
		openfile, err := os.OpenFile(clone.path+"/"+filename+"-"+environment+extension, os.O_RDONLY, 0600)
		if err != nil {
			continue
		}

		content, err := io.ReadAll(openfile)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))
		// string_data := string(content)

		c.Data(200, "application/json", content)
	}

}
