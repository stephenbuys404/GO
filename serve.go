-->main.go

package main
import (
        'fmt'
        'log'
        'net/http'
)
func formHandler(w http.ResponseWriter, r *http.Request) {
        if err := r.ParseForm(); err != nil {
                fmt.Fprintf(w, 'ParseForm() err: %v', err)
                return
        }
        fmt.Fprintf(w, 'POST request successful')
        name := r.FormValue('name')
        address := r.FormValue('address')
        fmt.Fprintf(w, 'Name = %s', name)
        fmt.Fprintf(w, 'Address = %s', address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != '/hello' {
                http.Error(w, '404 not found', http.StatusNotFound)
                return
        }
        if r.Method != 'GET' {
                http.Error(w, 'method is not supported', http.StatusNotFound)
                return
        }
        fmt.Fprintf(w, 'hello!')
}
func main() {
        fileServer := http.FileServer(http.Dir('./static'))
        http.Handle('/', fileServer)
        http.HandleFunc('/form', formHandler)
        http.HandleFunc('/hello', helloHandler)
        fmt.Printf('Starting server at port 8080')
        if err := http.ListenAndServe(':8080', nil); err != nil {
                log.Fatal(err)
        }
}

-->static/form.html

<!DOCTYPE html>
<html>
<head>
    <meta charset='UTF-8' />
    <title>Hello</title>
</head>
<body>
    <div>
        <form method='post' action='/form'>
        <label>Name</label>
        <input name='name' type='text' value='' />
        <label>Address</label>
        <input name='address' type='text' value='' />
        <input type='submit' value='submit' />
        </form>
    </div>
</body>
</html>


-->static/index.html

<html>
    <head><title>Hello</title></head>
    <body>
        <h2>Static website</h2>
    </body>
</html>