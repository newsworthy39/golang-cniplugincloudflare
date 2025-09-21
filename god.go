package main

import (
    "fmt"
    "strings"
)

type God struct {
    Commands map[string]Command
    Options map[string]string
}

func NewGod() *God {
    return &God{
        Commands: make(map[string]Command),
        Options: make(map[string]string),
    }
}

func (f *God) ValidateOptions(required []string) (string, bool) {
    for _, v := range required {
        key := strings.ToUpper(fmt.Sprintf("CNI_%s", v))
        _, ok := f.Options[key]
        if ok != true {
            return fmt.Sprintf("Key  %s does not exist.", key), true
        }
    }
    return "", false
}



func (g *God) RegisterCommand(name string, cmd Command) {
    g.Commands[name] = cmd
}

func (g *God) Execute(name string, input *Input) {
    if cmd, ok := g.Commands[name]; ok {
        cmd.Run(input, g.Options)
    } else {
        println("Unknown command:", name)
    }
}
