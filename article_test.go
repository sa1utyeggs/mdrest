package mdrest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadArticle(t *testing.T) {
	dir := "/Users/leenanxi/Documents/go/src/github.com/ti/mdrest/sample_docs/first dir/image.md"
	ar, err := ReadArticle("", dir, "https://static.lnx.cm/")
	if err != nil {
		log.Println(err)
	}

	_ = ar
}

func TestReadArticles(t *testing.T) {
	path := "C:\\Users\\ab875\\Desktop\\mdrest-master\\sample_docs\\YAML Article.md"
	ars, err := ReadArticles(path, "", true)
	if err != nil {
		log.Println(err)
	}
	ars.WriteFiles("C:\\Users\\ab875\\Desktop\\mdrest-master\\sample_docs\\", "json", true)

	siteMap := ars.GetSiteMap(2)

	jsb, _ := json.MarshalIndent(siteMap, "", "\t")

	log.Println(string(jsb))
}

func TestFrontMatter(t *testing.T) {
	fpath := "C:\\Users\\ab875\\Desktop\\mdrest-master\\sample_docs\\YAML Article.md"
	file, err := os.Open(fpath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	firstLine, _ := reader.ReadBytes(byte('\n'))
	fmt.Println(firstLine)

	matter, err := parseFrontMatter(reader)
	if err != nil {
		t.Error(err)
	}
	t.Log(matter)
}
