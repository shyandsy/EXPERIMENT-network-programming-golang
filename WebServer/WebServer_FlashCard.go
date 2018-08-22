/*
WebServer_FlashCard
*/
package main

type FlashCard struct {
	Simplified string
	English    string
	Dictionary *Dictionary
}

type FlashCards struct {
	Name      string
	CardOrder string
	ShowHalf  string
	Cards     []*FlashCard
}
