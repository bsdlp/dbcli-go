package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Program struct {
    Title    string `json:"title"`
    Id       int    `json:"id"`
    ImageUrl string `json:"image_url"`
}

type ProgramCollection []Program

type Workout struct {
    ImageUrl           string `json:"image_url"`
    WorkoutDescription string `json:"workout_description"`
    Title              string `json:"title"`
    ProgramIDs         []int  `json:"program_ids"`
    TrainerName        string `json:"trainer_name"`
}

type WorkoutCollection []Workout

func getPrograms() ProgramCollection {
    program_url := "http://dbios.herokuapp.com/programs"
    res, err := http.Get(program_url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    var result ProgramCollection
    if err := json.Unmarshal(body, &result); err != nil {
        panic(err)
    }
    return result
}

func main() {
    bleh := getPrograms()
    fmt.Println(bleh[0].Title)
}
