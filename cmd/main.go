package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Jorik887/bankAPI/cmd/api"
	"github.com/Jorik887/bankAPI/db"
	"github.com/Jorik887/bankAPI/types"
)

func seedAccount(store db.Storage, fname, lname, pw string) *types.Account {
	acc, err := types.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s db.Storage) {
	seedAccount(s, "jopa", "majopa", "testing1234")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the database")
		seedAccounts(store)

	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
