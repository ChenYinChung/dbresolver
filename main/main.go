package main

import (
	"dbresolver/handle"
	"fmt"
)

// Game is game
type Game struct {
	ID       int    `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	WinnerID int    `gorm:"type:int(5);" json:"winner_id,omitempty"`
	Location string `gorm:"type:varchar(20) NOT NULL;" json:"location,omitempty"`
	Prize    int    `gorm:"type:int(5);" json:"prize,omitempty"`
}

// Player is player
type Player struct {
	ID       int    `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Age      int    `gorm:"type:int(5);" json:"status,omitempty"`
	Username string `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Budget   int    `gorm:"type:int(5);" json:"budget,omitempty"`
}

func main() {
	dbClient := handle.DBClient{}

	err := dbClient.InitDB()
	if err != nil {
		panic(err)
	}

	player := Player{ID: 1}
	dbClient.Find(&player)

	fmt.Print(player)

	dbClient.Insert(&Game{Location: "Lion111", Prize: 1000})
	if dbClient.Err != nil {
		fmt.Println(dbClient.Err)
	}

	dbClient.Insert(&Player{Age: 10, Username: "Sammy", Budget: 1000})
	if dbClient.Err != nil {
		fmt.Println(dbClient.Err)
	}

	//
	//game := Game{
	//	WinnerID:1,
	//	Location:"TT",
	//	Prize:100,
	//}

	var games []Game

	dbClient.Retrieved(&games, 2, 3)
	for _, g := range games {
		fmt.Println(g)
	}

	fmt.Println("-------------------------")

	var players []Player

	dbClient.Retrieved(&players, 3, 1)
	for _, p := range players {
		fmt.Println(p)
	}

	fmt.Println("-------------------------")

}
