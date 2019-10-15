```
package main

import ("fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //return a response and an error, _ defines a variable we aren't using
	bytes, _ := ioutil.ReadAll(resp.Body) //after we get the data, we want to read the the bytes of data that are within the body
	stringBody := string(bytes) //convert the data bytes into strings
	fmt.Println(stringBody) //Print the internal body tags
	resp.Body.Close()
}
```
