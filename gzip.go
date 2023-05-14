package main
import (
    "compress/gzip"
    "os"
)
func main() {
    f, _ := os.Create("test.txt.gz")
    defer f.Close()
    gzWriter := gzip.NewWriter(f)
    content := "Hello World!"
    gzWriter.Write([]byte(content))
    gzWriter.Close()
}