package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Trinitui/awpost052"
)

var MIN1 = 0
var MAX1 = 26

func random01(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString01(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random01(MIN1, MAX1)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	awpost052.CHostname = "localhost"
	awpost052.CPort = 5433
	awpost052.CUsername = "postgres"
	awpost052.CPassword = "root"
	awpost052.CDatabase = "msds"

	data, err := awpost052.ListCourses()
	if err != nil {
		fmt.Println(err)
		//return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString01(5)
	random_coursename := "MSDS " + string(getString01(7))
	random_prereq := "MSDS_p " + string(getString01(2))

	t := awpost052.MSDSCourse{
		CID:     random_username,
		CNAME:   random_coursename,
		CPREREQ: random_prereq,
	}

	/*
		id := awpost052.AddCourse(t)
		if id == -1 {
			fmt.Println("There was an error adding Course", t.CNAME)
		}


			err = awpost052.DeleteCourse(id)
			if err != nil {
				fmt.Println(err)
			}

			// Trying to delete it again!
			err = awpost052.DeleteCourse(id)
			if err != nil {
				fmt.Println(err)
			}
	*/

	t = awpost052.MSDSCourse{
		CID:     random_username,
		CNAME:   random_coursename,
		CPREREQ: random_prereq,
	}

	fmt.Println(t)

	id := awpost052.AddCourse(t)
	if id == -1 {
		fmt.Println("There was an error adding course", t.CNAME)
	}

	err = awpost052.UpdateCourse(t)
	if err != nil {
		fmt.Println(err)
	}
}
