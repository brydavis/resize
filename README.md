# Resize

Quick image resizing

```go
package main

import (
	"fmt"
	"os"
	rs "github.com/brydavis/resize"
)

func main() {
	file, _ := os.Open("example.jpg")
	img, _ := rs.ResizePixels(file, 2000, 2000)
	err := rs.WriteToFile("new_example.jpg", img)
	if err != nil {
		fmt.Println(err)
	}
}
```

<br>
<br>

<hr>
<small>
<strong>MIT License &copy; 2015</strong>
</small>