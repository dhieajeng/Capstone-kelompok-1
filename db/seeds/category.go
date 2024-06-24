package seeds

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

func CreateCategorySeeds(db *gorm.DB) {
	data := []entity.Category{
		{Name: "Festival, Fair, Bazaar"},
		{Name: "Konser"},
		{Name: "Pertandingan"},
		{Name: "Exhibition, Expo, Pameran"},
		{Name: "Konferensi"},
		{Name: "Workshop"},
		{Name: "Pertunjukan"},
		{Name: "Atraksi, Theme Park"},
		{Name: "Akomodasi"},
		{Name: "Seminar, Talk Show"},
		{Name: "Social Gathering"},
		{Name: "Training, Sertifikasi, Ujian"},
		{Name: "Pensi, Event Sekolah, Kampus"},
		{Name: "Trip, Tur"},
		{Name: "Turnamen, Kompetisi"},
		{Name: "Lainnya"},
	}

	for _, category := range data {
		if err := db.Create(&category).Error; err != nil {
			fmt.Printf("Error when create category %s: %s\n", category.Name, err)
		}
	}
}
