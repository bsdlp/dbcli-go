package main

import (
    "github.com/codegangsta/cli"
    "os"
)

type Program struct {
    title     string
    id        int
    image_url string
}

type Workout struct {
    image_url           string
    workout_description string
    title               string
    program_ids         []int
    trainer_name        string
}

func main() {
    cli.NewApp().Run(os.Args)
}
