/*
WebServer_Server
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

// 词典
var d *Dictionary

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " :port\n")
		os.Exit(1)
	}

	port := os.Args[1]

	// 加载词库
	dictionaryPath := "cedict_ts.u8"
	d = new(Dictionary)
	d.Load(dictionaryPath)
	fmt.Println("Load dict ", len(d.Entries))

	http.HandleFunc("/", listFlashCards)
	fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
	http.Handle("/js/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)

	http.HandleFunc("/wordlook", lookupWord)
	http.HandleFunc("/flashcards.html", listFlashCards)
	http.HandleFunc("/flashcardSets", manageFlashCards)
	http.HandleFunc("/searchWord", searchWord)
	http.HandleFunc("/addWord", addWord)
	http.HandleFunc("/newFlashCardSet", newFlashCardSet)

	err := http.ListenAndServe(port, nil)
	checkError(err)
}

func indexPage(rw http.ResponseWriter, request *http.Request) {
	index, _ := ioutil.ReadFile("html/index.html")
	rw.Write([]byte(index))
}

type STD struct {
	Words *Dictionary
	Word  string
}

func lookupWord(rw http.ResponseWriter, request *http.Request) {
	word := request.FormValue("word")

	t := template.New("DictionaryEntry.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/DictionaryEntry.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if word != "" {
		words := d.LookupEnglish(word)
		t.Execute(rw, STD{Words: words, Word: word})
	} else {
		t.Execute(rw, nil)
	}
}

type DictPlus struct {
	Dict     *Dictionary
	Word     string
	CardName string
}

func searchWord(rw http.ResponseWriter, request *http.Request) {
	word := request.FormValue("word")
	searchType := request.FormValue("searchType")
	cardName := request.FormValue("cardname")

	var words *Dictionary
	var dp []DictPlus
	if searchType == "english" {
		words = d.LookupEnglish(word)
		d1 := DictPlus{Dict: words, Word: word, CardName: cardName}
		dp = make([]DictPlus, 1)
		dp[0] = d1
	} else {
		words = d.LookupPinyin(word)

		// 统计数量，分配内存
		numTrans := 0
		for _, entry := range words.Entries {
			numTrans += len(entry.Translations)
		}
		dp = make([]DictPlus, numTrans)

		// 开始逐个处理
		idx := 0
		for _, entry := range words.Entries {
			for _, trans := range entry.Translations {
				dict := new(Dictionary)
				dict.Entries = make([]*Entry, 1)
				dict.Entries[0] = entry
				dp[idx] = DictPlus{
					Dict:     dict,
					Word:     trans,
					CardName: cardName,
				}
				idx++
			}
		}
	}

	t := template.New("ChooseDictionaryEntry.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/ChooseDictionaryEntry.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(rw, dp)
}

func newFlashCardSet(rw http.ResponseWriter, request *http.Request) {
	defer http.Redirect(rw, request, "http://flashcards.html", 200)

	newSet := request.FormValue("NewFlashcard")
	fmt.Println("New cards", newSet)

	b, err := regexp.Match("[/$~]", []byte(newSet))
	if err != nil {
		return
	}
	if b {
		fmt.Println("No good string")
		return
	}
	//flashcards.NewFlashCardSet(newSet)
	return
}

func addWord(rw http.ResponseWriter, request *http.Request) {
	url := request.URL
	fmt.Println("url: ", url.String())
	fmt.Println("query: ", url.RawQuery)

	word := request.FormValue("word")
	cardName := request.FormValue("cardname")
	simplified := request.FormValue("simplified")
	pinyin := request.FormValue("pinyin")
	traditional := request.FormValue("traditional")
	translations := request.FormValue("translations")

	fmt.Println("word is: ", word, ", card is ", cardName,
		", simplified is ", simplified, ", pinyin is ", pinyin,
		"trad is ", traditional, ", trans is ", translations)
	// flashcards.AddFlashEntry(cardName, word, pinyin, simplified,
	//traditional, translations)

	//add another card?
	//addFlashCards(rw, cardName)

	return
}

func listFlashCards(rw http.ResponseWriter, request *http.Request) {
	//flashCardsName :=
	return
}

func manageFlashCards(rw http.ResponseWriter, req *http.Request) {
	return
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	return
}

func listWords(rw http.ResponseWriter, cardname string) {
	return
}

func addFlashCards(rw http.ResponseWriter, cardname string) {
	return
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
