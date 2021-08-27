// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Civilization struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Expansion         string   `json:"expansion"`
	ArmyType          string   `json:"army_type"`
	UniqueUnit        []string `json:"unique_unit"`
	UniqueTech        []string `json:"unique_tech"`
	TeamBonus         string   `json:"team_bonus"`
	CivilizationBonus []string `json:"civilization_bonus"`
}

type Formula struct {
	Id              int
	Name            string
	FullInformation string
}

func (f Formula) Run(writer io.Writer) {
	var result string

	result += fmt.Sprintf("AoE API!\n")

	result += color.FgBlue.Render(fmt.Sprintf("The id about your civilization request is %d. \n", f.Id))

	result += color.FgGreen.Render(fmt.Sprintf("You requested for informations about civilization %s.\n", f.Name))

	result += color.FgBlue.Render(fmt.Sprintf("The full informations is listed bellow:\n%v", f.FullInformation))

	result += fmt.Sprintf("::output key=%s\n", f.Name)

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
