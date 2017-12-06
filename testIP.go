package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func connTest(dest, src string) {

    file, err := os.Open(src)
    check(err)

    f, err := os.Create(dest)
    check(err)

    defer file.Close()
    defer f.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ip := scanner.Text()
        addr, err := net.LookupAddr(ip)
        if err != nil {
            out, _ := exec.Command("ping", ip, "-w", "2").Output()
            if strings.Contains(string(out), "100% straty") {
                fmt.Fprintln(f, fmt.Sprintf("%s %s", ip, "- TANGO DOWN"))
            } else {
                fmt.Fprintln(f, fmt.Sprintf("%s %s", ip, "- Work only ping"))
            }
        } else {
            fmt.Fprintln(f, fmt.Sprintf("%s - %s", ip, addr))
        }
    }

}

func main() {

    var src = "file.txt"
    var dest = "wynik.txt"
    connTest(dest, src)

}
