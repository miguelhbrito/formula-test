// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"encoding/json"
	"fmt"
	"formula/pkg/formula"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	input1 := os.Getenv("RIT_INPUT_TEXT")

	url := "https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/" + input1
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error to get infos about civilization %s, error:\n %s", input1, err))
		return
	}
	defer resp.Body.Close()

	civilization := formula.Civilization{}

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &civilization)

	marshallCiv, error := json.MarshalIndent(civilization, "", "\t")
	if error != nil {
		fmt.Println("JSON parse error: ", error)
		return
	}

	formula.Formula{
		Id:              civilization.ID,
		Name:            civilization.Name,
		FullInformation: string(marshallCiv),
	}.Run(os.Stdout)

}
