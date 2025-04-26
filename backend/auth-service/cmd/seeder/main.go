package main

import (
	"auth-service/config"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tanimutomo/sqlfile"
)

/// How re-build this script?
/// run: ./script/build.sh

const (
	MASTER_DIR string = "./database/seeders/master"
	TEST_DIR   string = "./database/seeders/test"
)

// / reference: https://qiita.com/tanimutomo/items/afd2503f7563555c27d6
func main() {
	// Get a database handler
	db, err := config.Initialize()
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Initialize SqlFile
	s := sqlfile.New()

	// Load Data Master
	runSeeder(s, sqlDB, MASTER_DIR)

	// Load Data Test
	runSeeder(s, sqlDB, TEST_DIR)

	fmt.Println("FINISHED")
}

func runSeeder(s *sqlfile.SqlFile, sqlDB *sql.DB, dir string) {
	fmt.Printf("Load dir %s\n", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for idx, file := range files {
		// Skip directory & non `*.sql` file
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		fmt.Printf("%d- Run %s\n", idx, file.Name())

		s := sqlfile.New()
		err := s.File(dir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		// Execute the stored queries
		res, err := s.Exec(sqlDB)
		if err != nil {
			panic(err)
		}

		// Print Result
		fmt.Printf("Results:\n")
		for idx, result := range res {
			rowAffected, _ := result.RowsAffected()
			fmt.Printf("%d - Row Affected: %d\n", idx, rowAffected)
		}
	}

	// err := s.Directory(dir)
	// if err != nil {
	// 	panic(err)
	// }

	// // Execute the stored queries
	// fmt.Printf("Exec %s \n", dir)
	// res, err := s.Exec(sqlDB)
	// if err != nil {
	// 	panic(err)
	// }

	// // Print Result
	// fmt.Printf("Results:\n")
	// for idx, result := range res {
	// 	rowAffected, _ := result.RowsAffected()
	// 	fmt.Printf("%d - Row Affected: %d\n", idx, rowAffected)
	// }

}
