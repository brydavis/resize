package main

func main() {
	// file, _ := os.Open("gophermega.jpg")
	// img, _ := ResizePixels(file, 2000, 2000)
	// WriteToFile("mega.jpg", img)

	ResizeFolderPixels(".", 100, 200, "sm")
}
