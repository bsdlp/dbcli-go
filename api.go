package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const apiEndpoint = "http://dbios.herokuapp.com/"

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

func getJawn(thingType string) []byte {
    apiURL := apiEndpoint + thingType
    res, err := http.Get(apiURL)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    return body
}

func parsePrograms() ProgramCollection {
    response := getJawn("programs")
    var result ProgramCollection
    if err := json.Unmarshal(response, &result); err != nil {
        panic(err)
    }

    return result
}

func parseWorkouts() WorkoutCollection {
    response := getJawn("workouts")
    var result WorkoutCollection
    if err := json.Unmarshal(response, &result); err != nil {
        panic(err)
    }

    return result
}

func listPrograms() {
    programs := parsePrograms()
    for i := range programs {
        fmt.Println(programs[i].Title, programs[i].Id, programs[i].ImageUrl)
    }
}

func main() {
    listPrograms()
}
