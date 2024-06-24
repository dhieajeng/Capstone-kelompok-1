package seeds

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

func CreateLocationSeeds(db *gorm.DB) {
	locations := []entity.Location{
		{
			Name: "Bandung",
		},
		{
			Name: "Daerah Khusus Jakarta",
		},
		{
			Name: "Bali",
		},
		{
			Name: "DI Yogyakarta",
		},
		{
			Name: "Surabaya",
		},
		{
			Name: "PIK 2",
		},
		{
			Name: "Kebun Raya Bogor",
		},
	}

	for _, location := range locations {
		if err := db.Create(&location).Error; err != nil {
			fmt.Printf("Error when create location %s: %s\n", location.Name, err)
		}
	}
}
