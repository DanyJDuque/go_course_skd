package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/DanyJDuque/go_course_skd/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("ebdfc1f2-0458-46c1-b4db-be34d6ebda69")
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
