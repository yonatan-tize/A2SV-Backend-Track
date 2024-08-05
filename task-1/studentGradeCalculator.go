package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define struct that represents a subject name and grade
type SubjectGrade struct {
	subjectName string
	grade       float32
}

// Slice to hold students' subjects and grades
var studentGrade = []SubjectGrade{}

// Function to take user input for the number of subjects and their grades
func userGradesPrompt() {

	reader := bufio.NewReader(os.Stdin)

	var totalNumSubjects int
	var err error

	// Prompt the user until they enter a valid number of subjects
	for {
		fmt.Print("Enter the number of subjects you have taken: ")
		val, _ := reader.ReadString('\n')
		val = strings.TrimSpace(val)
		totalNumSubjects, err = strconv.Atoi(val)
		if err == nil {
			break
		}
		fmt.Println("Invalid number, please enter a valid integer.")
	}

	// Iterate to accept user input for each subject and its grade
	for i := 0; i < totalNumSubjects; i++ {

		// Prompt the user until a valid subject name is entered
		var subjName string
		for {
			fmt.Print("Enter the name of the subject: ")
			subName, _ := reader.ReadString('\n')
			subjName = strings.TrimSpace(subName)
			if subjName != "" {
				break
			}
			fmt.Println("Invalid subject name.")
		}

		// Prompt the user until a valid grade is entered
		var subjGrade float32
		for {
			fmt.Print("Enter the grade you got: ")
			grade, _ := reader.ReadString('\n')
			grade = strings.TrimSpace(grade)
			var grade64 float64 // Convert to float64 since ParseFloat returns float64
			grade64, err = strconv.ParseFloat(grade, 32)
			subjGrade = float32(grade64)

			// Check if the grade is valid
			if err == nil && 0 <= subjGrade && subjGrade <= 100 {
				break
			}
			if err != nil {
				fmt.Println("Invalid number, please enter a valid float.")
			} else {
				fmt.Println("Not a valid grade.")
			}
		}

		// Create a new SubjectGrade to be appended to the studentGrade slice
		newSubject := SubjectGrade{
			subjectName: subjName,
			grade:       subjGrade,
		}

		studentGrade = append(studentGrade, newSubject)
	}
}

func main() {
	// Prompt the name of the student
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	fmt.Printf("Name: %s\n", studentName)

	// Call the function to take user input for grades
	userGradesPrompt()

	// Print the number of subjects
	fmt.Printf("Number of subjects: %d\n", len(studentGrade))
	fmt.Println() // Create a new line for readability

	var totalGrade float32
	// Iterate over the slice to print each subject-grade pair and calculate the total grade
	for _, subjectGrade := range studentGrade {
		fmt.Printf("Subject: %s, Grade: %.2f\n", subjectGrade.subjectName, subjectGrade.grade)
		totalGrade += subjectGrade.grade
	}
	// Calculate and print the total and average grades
	average := totalGrade / float32(len(studentGrade))
	fmt.Printf("Your total grade is: %.2f, and your average grade is: %.2f\n", totalGrade, average)
}
