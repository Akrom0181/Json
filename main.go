package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "math/rand"
  "os"
  "strconv"
)

type Student struct {
  Name     string `json:"name"`
  Age      int    `json:"age"`
  Course   int    `json:"course"`
  Contract int    `json:"contract"`
}

type Course struct {
  Count    int
  Contract int
}

func main() {

  	file, err := os.Open("students.json")
  	if err != nil {
    	fmt.Println("error:", err)
    	return
  	}
  	defer file.Close()
  	var students []Student

  	decoder := json.NewDecoder(file)
  	if err := decoder.Decode(&students); err != nil {
    	fmt.Println("Error decoding JSON: ", err)
    	return
  	}
  	fmt.Println(students)

  	courseMap := make(map[int]Course)
  	path := "./files/"
  	filesPaths := []string{}

  	for _, student := range students {
    	fileName := path + student.Name + ".txt"
    	newFileName := fileName
    if checkFileExists(fileName) {
      	randNumb := strconv.Itoa(rand.Intn(10))
      	newFileName = fmt.Sprintf("%v.txt", path+student.Name+randNumb)
    }
    filesPaths = append(filesPaths, newFileName)
    _, err := os.Create(newFileName)
    if err != nil {
      	fmt.Println("error while creating file: ", err)
    }
    currentvalue := courseMap[student.Course]
    courseMap[student.Course] = Course{
      Count:    currentvalue.Count + 1,
      Contract: currentvalue.Contract + student.Contract,
    }
  }
  fmt.Println("courses", courseMap)

  var choice int
  fmt.Println("if you want to delete files, select 1 otherwise 0")
  fmt.Scan(&choice)
  if choice == 1 {
    for _, path := range filesPaths {
      	err := os.Remove(path)
      	if err != nil {
        	fmt.Println("error while removing files: ", err)
        	return
      	}
    }
    fmt.Println("All files deleted")
  } else {
    fmt.Println("Goodbye")
  }

}

func checkFileExists(filePath string) bool {
  	_, error := os.Stat(filePath)
  	return !errors.Is(error, os.ErrNotExist)
}