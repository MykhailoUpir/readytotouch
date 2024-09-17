package main

import (
	"fmt"
	"go/format"
	"os"

	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/protos/organizers"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

func main() {
	var (
		maxID = int64(0)
		pairs = make([]*dev.CompanyCodePair, len(db.Companies()))
	)

	// Assert that all company aliases are present in the map
	for alias := range organizers.CompanyStartupMap {
		if _, ok := organizers.CompanyAliasMap[alias]; !ok {
			panic(fmt.Sprintf("Company alias not found: %s", alias))
		}
	}

	for i, company := range db.Companies() {
		pair := &dev.CompanyCodePair{
			ID:    organizers.CompanyAliasMap[company.LinkedInProfile.Alias],
			Name:  company.LinkedInProfile.Name,
			Alias: company.LinkedInProfile.Alias,
		}

		pairs[i] = pair
		maxID = max(maxID, pair.ID)
	}

	for _, pair := range pairs {
		if pair.ID == 0 {
			maxID++
			pair.ID = maxID
		}
	}

	output, err := format.Source([]byte(dev.CompanyCode(pairs)))
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./internal/protos/organizers/company_code.go", output, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Company code generated successfully")
	fmt.Printf("Max ID: %d\n", maxID)

	// json.NewEncoder(os.Stdout).Encode(pairs)
}