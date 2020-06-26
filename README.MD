# go-file-generate
## Description
Generates a file on provided location of wanted size in bytes

## Usage
``` golang
import fg "github.com/alexalexiuc/go-file-generate"

func main() {
	err := fg.GenerateFile("./file.txt", 1024)
	if err != nil {
		panic(err)
	}
}
```
