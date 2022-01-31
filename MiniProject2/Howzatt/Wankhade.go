package howzatt

import (
	"fmt"
	"math/rand"
	"time"
)

func Wankhade() {
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Println("Welcome to Wankhade! ")

	var wickets int = 4
	// fmt.Println("Final score of India", "<", runs, "/", wickets, ">")

	//Print Names and Scores of all Indian batsmen, in format (Name --> Score)
	var arr [11]string
	arr[0] = "MSDhoni-Captain"
	arr[1] = "VirenderSehwag(Vice-captain)"
	arr[2] = "SachinTendulkar"
	arr[3] = "YuvrajSingh"
	arr[4] = "GauitamGambhir"
	arr[5] = "SureshRaina"
	arr[6] = "ZaheerKhan"
	arr[7] = "MunafPatel"
	arr[8] = "AshishNehra"
	arr[9] = "ViratKohli"
	arr[10] = "HarbhajanSingh"

	fmt.Println("Players in the team are:\n", arr)
	fmt.Println("Names and Scores of all Indian batsmen, in format (Name --> Score)")
	var score, i int
	var runs int = 0
	rand.Seed(time.Now().UnixNano())
	for i = 0; i < 10; i++ {
		score = rand.Intn(50)
		fmt.Println(arr[i], "-->", score)
		runs = score + runs
	}
	fmt.Println(runs)
	fmt.Println("Final score of India", "<", runs, "/", wickets, ">")

	//Names and Runs conceded by SL bowlers in format(Name --> Runs conceded)
	var arr1 [6]string
	arr1[0] = "Lasith Malinga"
	arr1[1] = "Nuwan Kulasekara"
	arr1[2] = "Thisara Perera"
	arr1[3] = "Suraj Randiv"
	arr1[4] = "Tillakaratne Dilshan"
	arr1[5] = "Muttiah Muralitharan"

	fmt.Println(arr1)

	fmt.Println("Name ---> RunsConceeded by srilanka")
	var score1, j int
	var runsConceeded int = 0
	rand.Seed(time.Now().UnixNano())
	for j = 0; j < 6; j++ {
		score1 = rand.Intn(50)
		fmt.Println(arr1[j], "-->", runsConceeded)
		runsConceeded = score1 + runsConceeded
	}
	fmt.Println(runsConceeded)

	//Print Names and Scores of all SL batsmen, in format (Name --> Score)

	var arrsl [11]string
	arrsl[0] = "Upul Tharanga"
	arrsl[1] = "Tillakaratne Dilshan"
	arrsl[2] = "Kumar Sangakkara"
	arrsl[3] = "Mahela Jayawardene"
	arrsl[4] = "Thilan Samaraweera"
	arrsl[5] = "Chamara Kapugedera"
	arrsl[6] = "Nuwan Kulasekara"
	arrsl[7] = "Thisara Perera"
	arrsl[8] = "Lasith Malinga"
	arrsl[9] = "Suraj Randiv"
	arrsl[10] = "Muttiah Muralitharan"

	fmt.Println("Players in the team are:\n", arrsl)
	fmt.Println("Names and Scores of all srilankan batsmen, in format (Name --> Score)")
	var scoresl, k int
	var runssl int = 0
	rand.Seed(time.Now().UnixNano())
	for k = 0; k < 10; k++ {
		scoresl = rand.Intn(50)
		fmt.Println(arrsl[k], "-->", scoresl)
		runssl = scoresl + runssl
	}
	fmt.Println(runssl)

	//Print Names and Runs conceded by Indian bowlers in format(Name --> Runs conceded)

	fmt.Println("Names and Runs conceded by Indian bowlers in format(Name --> Runs conceded)")
	var arr2 [7]string
	arr2[0] = "Zaheer Khan"
	arr2[1] = "Sreesanth"
	arr2[2] = "Munaf Patel"
	arr2[3] = "Harbhajan Singh"
	arr2[4] = "Yuvraj Singh"
	arr2[5] = "Sachin Tendulkar"
	arr2[6] = "Virat Kohli"

	fmt.Println(arr2)

	fmt.Println("Name ---> RunsConceeded by India")
	var score2, m int
	var runsConceededIndia int = 0
	rand.Seed(time.Now().UnixNano())
	for m = 0; m < 6; m++ {
		score2 = rand.Intn(50)
		fmt.Println(arr2[m], "-->", runsConceededIndia)
		runsConceededIndia = score2 + runsConceededIndia
	}
	fmt.Println(runsConceededIndia)

	if runs > runssl {
		fmt.Println("India Won")
	} else {
		fmt.Println("Srilanka won")
	}
}
