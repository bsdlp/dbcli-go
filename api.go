package main

import (
    "github.com/codegangsta/cli"
    "os"
)

type Program struct {
    Title    string `json:"title"`
    Id       int    `json:"id"`
    ImageUrl string `json:"image_url"`
}

type ProgramCollection struct {
    Programs []Program
}

type Workout struct {
    ImageUrl           string `json:"image_url"`
    WorkoutDescription string `json:"workout_description"`
    Title              string `json:"title"`
    ProgramIDs         []int  `json:"program_ids"`
    TrainerName        string `json:"trainer_name"`
}

type WorkoutCollection struct {
    Workouts []Workout
}

func main() {
    cli.NewApp().Run(os.Args)
}
