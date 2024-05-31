package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxStudents    = 30
	totalStudents  = 100
	maxReadingTime = 4
)

// Student represents a student in the library
type Student struct {
	id   int
	time int
}

func main() {
	// Create a channel to store students
	students := make(chan Student, maxStudents)
	// Create a wait group to wait for all students to finish
	var wg sync.WaitGroup

	// Create a goroutine for each student
	for i := 0; i < totalStudents; i++ {
		time.Sleep(time.Second)
		duration := rand.Intn(maxReadingTime) + 1
		student := Student{
			id:   i,
			time: duration,
		}
		// Add the student to the wait group
		wg.Add(1)
		go func(s Student) {
			defer wg.Done()
			students <- s
			fmt.Printf("Time %d: student %d starts reading at the library\n", time.Now().Second(), s.id)
			time.Sleep(time.Second * time.Duration(s.time))
			fmt.Printf("Time %d: student %d leave the library with %d hours\n", time.Now().Second(), s.id, s.time)
			<-students
		}(student)
	}

	go func() {
		wg.Wait()
		close(students)
	}()

	fmt.Print("No more students in the library\n")

}
