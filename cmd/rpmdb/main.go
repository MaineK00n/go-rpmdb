package main

import (
	"fmt"
	"log"

	rpmdb "github.com/knqyf263/go-rpmdb/pkg"
)

func main() {
	db, err := rpmdb.Open("./Packages.db")
	if err != nil {
		db, err = rpmdb.Open("./Packages")
	}

	if err != nil {
		log.Fatal(err)
	}
	pkgList, err := db.ListPackages()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Packages:")
	for _, pkg := range pkgList {
		// Suppress output
		pkg.BaseNames = nil
		pkg.DirIndexes = nil
		pkg.DirNames = nil

		fmt.Printf("\t%+v\n", *pkg)
	}
	fmt.Printf("[Total Packages: %d]\n", len(pkgList))
}
