package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)
func main() {
    client := &http.Client{}
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        url := "https://techoverflow.net" + r.URL.Path
        req, _ := http.NewRequest(r.Method, url, r.Body)
        req.Header = r.Header
        resp, _ := client.Do(req)
        defer resp.Body.Close()
        header := w.Header()
        for name, headers := range resp.Header {
            for _, hdr := range headers {
                header.Add(name, hdr)
            }
        }
        body, _ := ioutil.ReadAll(resp.Body)
        w.WriteHeader(resp.StatusCode)
        w.Write(body)
    })
    fmt.Println("Listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}