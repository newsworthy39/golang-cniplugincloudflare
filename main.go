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
    
    // Try JSON stdin
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        input = ReadJSONInput() // returns Input-struct.
    } 
        
    // Read Environments variables.
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        if strings.HasPrefix(pair[0], "CNI") {
            god.Options[pair[0]] = pair [1]
        }
    }

    // These are required arguments, to be found in the environments.
    required := []string {
        "command",
        "netns",
        "ifname",
    }

    // Make sure, to test required options.
    msg, notok := god.ValidateOptions( required) 
    if ( notok ) {
        panic(msg)
    }
    
    // And finally, execute it
    god.Execute(god.Options["CNI_COMMAND"], input)

    // Report the original json back
    fmt.Println(string(input.raw))
}
