package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/DanyJDuque/go_course_skd/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("0255f4d6-7db5-405b-8d4c-1fdb9b1cbd37")
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
