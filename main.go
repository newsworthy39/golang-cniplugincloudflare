package main

import (
    "os"
    "strings"
    "fmt"
)

func main() {
    god := NewGod()
    god.RegisterCommand("ADD", AddCommand{})
    god.RegisterCommand("DEL", DelCommand{})

    var input *Input
    if len(os.Args) > 1 {
        // Command-line mode
        input = &Input{}
    } else {
        // Try JSON stdin
        stat, _ := os.Stdin.Stat()
        if (stat.Mode() & os.ModeCharDevice) == 0 {
            input = ReadJSONInput() // returns Input-struct.
        } 
    }
    
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        if strings.HasPrefix(pair[0], "CNI") {
            god.Options[pair[0]] = pair [1]
        }
    }
    
    god.Execute(god.Options["CNI_COMMAND"], input)

    // Report the original json back
    fmt.Println(string(input.raw))
}
