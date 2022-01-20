package hogwarts

import (
	"fmt"
	"math/rand"
	"time"
)

func HarryPotter() {
	fmt.Println("Introduction lunch and starting segregating students in 4 groups hosted by Prof. Dumbledore ")
	var students = make([]int, 160) //Created a slice of length 160 and as we didn't declare the capacity it is by default 160
	var gryffindor = students[0:40]
	var slytherin = students[40:80]
	var ravenclaw = students[80:120]
	var hufflepuff = students[120:160] //Here we have created each slice of length 40 i.e. I have created gryffindor, slytherin, ravenclaw, hufflepuff groups each of 40 student
	rand.Seed(time.Now().UnixNano())   //This is to generate random numbers for an array or slice
	for i := 0; i < 40; i++ {          //Loop until all the 40 students are given some random numbers
		gryffindor[i] = rand.Intn(100)
		slytherin[i] = rand.Intn(100)
		ravenclaw[i] = rand.Intn(100)
		hufflepuff[i] = rand.Intn(100) //Each group of students are assigned with some random number

	}
	fmt.Println("List of points scored by each student in  group Gryffindor:", gryffindor) //Random numbers generated for each student in gryffindor grou
	fmt.Println("List of points scored by each student in  group Slytherin:", slytherin)   //Random numbers generated for each student in slytherin group
	fmt.Println("List of points scored by each student in  group Ravenclaw:", ravenclaw)   //Random numbers generated for each student in ravenclaw group
	fmt.Println("List of points scored by each student in  group Hufflepuff:", hufflepuff) //Random numbers generated for each student in hufflepuff group

	//This is to pint the total number of points scored by each group in total
	var result int
	var result1, result2, result3 int
	for i := 0; i < 40; i++ {
		result = result + gryffindor[i]
		result1 = result1 + slytherin[i]
		result2 = result2 + ravenclaw[i]
		result3 = result3 + hufflepuff[i]
	}
	fmt.Println("Total points scored by group Gryffindor:", result)
	fmt.Println("Total points scored by group Slytherin:", result1)
	fmt.Println("Total points scored by group Ravenclaw:", result2)
	fmt.Println("Total points scored by group Hufflepuff:", result3)

	//Due to a mis-use of magic trick by few students of Slytherin, Prof. Dumbledore decides to cut 2 points for students with id 10, 25, 29
	slytherin[9] = slytherin[9] - 2
	slytherin[24] = slytherin[24] - 2
	slytherin[28] = slytherin[28] - 2
	//This is to give reward points de to  an excellent work of Hermoine (id:9), Harry potter(id:27) and Ron Weasley(id:15)
	gryffindor[8] = gryffindor[8] + 4 //we change value of id 9 i.e in index 8
	gryffindor[26] = gryffindor[26] + 4
	gryffindor[14] = gryffindor[14] + 4

	//One point is deducted for eachstudent in Ravenclaw because it lost the “quiddich” game(game played with brooms)
	for i := 0; i < len(ravenclaw); i++ {
		ravenclaw[i] = ravenclaw[i] - 1
	}
	// fmt.Println(ravenclaw)

	//Two reward points are given to each in Hufflepuff because it stood 3rd place in “quiddich” game
	for i := 0; i < len(hufflepuff); i++ {
		hufflepuff[i] = hufflepuff[i] + 2
	}
	// fmt.Println(hufflepuff)

	//This is to print the list of scores of each student in each group...
	fmt.Println("List of points scored by each student in group Gryffindor:", gryffindor)
	fmt.Println("List ofpoints scored by each student in group Slytherin:", slytherin)
	fmt.Println("List of points scored by each student n group Ravenclaw:", ravenclaw)
	fmt.Println("List of poins scored by each student in group Hufflepuff:", hufflepuff)

	//This is to print the total points scored by each group
	var changeResult int
	var changeResult1, changeResult2, changeResult3 int
	for i := 0; i < len(gryffindor); i++ { //length is taken as len(gryffindor) because length of all the groups are same
		changeResult = changeResult + gryffindor[i]
		changeResult1 = changeResult1 + slytherin[i]
		changeResult2 = changeResult2 + ravenclaw[i]
		changeResult3 = changeResult3 + hufflepuff[i]
	}
	fmt.Println("Total points scored by group Gryffindor", changeResult)
	fmt.Println("Total pointsscored by group Slytherin", changeResult1)
	fmt.Println("Total points scored by group Ravenclaw", changeResult2)
	fmt.Println("Total points scored by group Hufflepuff", changeResult3)

	//This is to print the average points scored by group gryffindor
	var AveragePoints float64
	AveragePoints = float64(changeResult) / float64(len(gryffindor))
	fmt.Println("Average Point scored by gryffindor:", AveragePoints)

	//This is to print the average points scored by group slytherin
	var AveragePoints1 float64
	AveragePoints1 = float64(changeResult1) / float64(len(slytherin))
	fmt.Println("Average Points scored by slytherin:", AveragePoints1)

	//This is to print the average points scored by group ravenclaw
	var AveragePoints2 float64
	AveragePoints2 = float64(changeResult2) / float64(len(ravenclaw))
	fmt.Println("Average Points scored by ravenclaw:", AveragePoints2)

	//This is to print the average points scored by group hufflepuff
	var AveragePoints3 float64
	AveragePoints3 = float64(changeResult3) / float64(len(hufflepuff))
	fmt.Println("Average Points scored by hufflepuff:", AveragePoints3)

}
