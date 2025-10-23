package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/DanyJDuque/go_course_skd/course"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("4d43ecf8-a4a8-4185-a8db-cd40f3352bc2")
	fmt.Println("Course:", course)
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println(course)
}
