package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	listPullsURL      = "https://api.github.com/repos/cosmos/launch/pulls"
	commentURLFmt     = "https://api.github.com/repos/cosmos/launch/issues/%d/comments"
	contentPathURLFmt = "https://api.github.com/repos/%s/contents/gentx"
)

var (
	accessToken = fmt.Sprintf("access_token=%s", token)
	token       = os.Getenv("GITHUB_API_TOKEN")
)

type pullsResponse struct {
	URL    string `json:"url"`
	Number int    `json:"number"`
}

type pullResponse struct {
	Head headType `json:"head"`
}

type headType struct {
	Ref  string   `json:"ref"`
	Repo repoType `json:"repo"`
}

type repoType struct {
	Name string `json:"full_name"`
}

type contentsResponse struct {
	Name string `json:"name"`
}

type contentResponse struct {
	Content string `json:"content"`
}

type commentResponse struct {
	User userType `json:"user"`
	Body string   `json:"body"`
}

type userType struct {
	Login string `json:"login"`
}

type genTx struct {
	Value genTxValue `json:"value"`
}

type genTxValue struct {
	Msg []genTxMsg `json:"msg"`
}

type genTxMsg struct {
	Value genTxMsgValue `json:"value"`
}

type genTxMsgValue struct {
	Description description `json:"description"`
	Value       amount      `json:"value"`
}

type description struct {
	Moniker string `json:"moniker"`
}

type amount struct {
	Amount string `json:"amount"`
}

func getURL(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return bz
}

func main() {
	if token == "" {
		fmt.Println("Please set GITHUB_API_TOKEN")
		os.Exit(1)
	}

	url := listPullsURL + "?" + accessToken + "&per_page=100"
	bz := getURL(url)
	var pulls []pullsResponse
	err := json.Unmarshal(bz, &pulls)
	if err != nil {
		fmt.Println(url)
		fmt.Println(string(bz))
		panic(err)
	}
	fmt.Println("PULLS", len(pulls))

	validators := make(map[string]float64)
	var total float64

	for _, thisPull := range pulls {
		url := thisPull.URL + "?" + accessToken
		bz := getURL(url)
		var pull pullResponse
		err = json.Unmarshal(bz, &pull)
		if err != nil {
			fmt.Println(url)
			fmt.Println(string(bz))
			panic(err)
		}
		repoName := pull.Head.Repo.Name
		fmt.Println(repoName)
		branch := pull.Head.Ref

		url = fmt.Sprintf(contentPathURLFmt+"?ref=%s&%s", repoName, branch, accessToken)
		bz = getURL(url)
		var files []contentsResponse
		err = json.Unmarshal(bz, &files)
		if err != nil {
			fmt.Println(url)
			fmt.Println(string(bz))
			panic(err)
		}

		if len(files) != 2 {
			fmt.Println("INVALID NUMBER OF FILES", len(files), repoName)
			continue
		}

		var name string
		for _, f := range files {
			if f.Name == "README.md" {
				continue
			}
			name = f.Name
		}
		url = fmt.Sprintf(contentPathURLFmt+"/"+name+"?ref=%s&%s", repoName, branch, accessToken)
		bz = getURL(url)
		var content contentResponse
		err = json.Unmarshal(bz, &content)
		if err != nil {
			fmt.Println(url)
			fmt.Println(string(bz))
			panic(err)
		}

		bz, err = base64.StdEncoding.DecodeString(content.Content)
		if err != nil {
			fmt.Println(string(content.Content))
			panic(err)
		}

		var thisGenTx genTx
		err = json.Unmarshal(bz, &thisGenTx)
		if err != nil {
			fmt.Printf("INVALID GENTX %s, %s", repoName, name) //string(bz))
			continue                                           // panic(err)
		}
		msg := thisGenTx.Value.Msg[0]
		name = msg.Value.Description.Moniker
		amount := msg.Value.Value.Amount
		uatoms, err := strconv.Atoi(amount)
		if err != nil {
			fmt.Println(amount)
			panic(err)
		}
		atoms := float64(uatoms) / 1000000

		if atoms > 100000 {
			fmt.Printf("INVALID GENTX %s, %s", repoName, name)
			continue
		}
		fmt.Println("...", name, atoms)
		validators[name] = atoms
		total += atoms
	}

	bz, err = json.MarshalIndent(validators, "", "  ")
	if err != nil {
		fmt.Println(string(bz))
		panic(err)
	}
	fmt.Println(string(bz))
	fmt.Println("Validators", len(validators))
	fmt.Println("TOTAL", total)

}
