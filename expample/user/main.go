package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/DanyJDuque/go_course_skd/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("ffbc29df-a7a3-4094-9372-b59352d6eae0")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
