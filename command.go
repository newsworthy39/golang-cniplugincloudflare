package main

type Command interface {
    Run(input *Input, options map[string]string)
}