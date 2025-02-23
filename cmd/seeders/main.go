package main

import (
	"fmt"

	"github.com/adislice/go-project-structure/config"
	"github.com/adislice/go-project-structure/internal/database"
	"github.com/adislice/go-project-structure/internal/database/seeders"
)

func main() {
	config.LoadConfig()
	db := database.ConnectDB()
	fmt.Println("Seeding started")

	// panggil function seeder di sini
	seeders.RoleSeeder(db)

	fmt.Println("Seeding completed")
}
