import (
    "io/ioutil"
    "net/http"
)
func GetExternalIPv4Address() (string, error) {
    resp, err := http.Get("https://api4.ipify.org")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return string(body), nil
}
func GetExternalIPv6Address() (string, error) {
    resp, err := http.Get("https://api6.ipify.org")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return string(body), nil
}