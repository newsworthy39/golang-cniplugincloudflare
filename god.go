package main

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
