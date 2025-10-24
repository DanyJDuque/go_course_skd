package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/DanyJDuque/go_course_skd/course"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("807e78d4-36a9-4bed-94ae-4fd81aa6c4e8")
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
