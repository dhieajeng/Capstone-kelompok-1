package seeds

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
)

func CreateTopicSeeds(db *gorm.DB) {
	data := []entity.Topic{
		{Name: "Anak, Keluarga"},
		{Name: "Bisnis"},
		{Name: "Desain, Foto, Video"},
		{Name: "Fashion, Kecantikan"},
		{Name: "Hobi, Kerajinan Tangan"},
		{Name: "Investasi, Saham"},
		{Name: "Karir, Pengembangan Diri"},
		{Name: "Keagamaan"},
		{Name: "Kesehatan, Kebugaran"},
		{Name: "Keuangan, Finansial"},
		{Name: "Lingkungan Hidup"},
		{Name: "Makanan, Minuman"},
		{Name: "Marketing"},
		{Name: "Musik"},
		{Name: "Olahraga"},
		{Name: "Otomotif"},
		{Name: "Sains, Teknologi"},
		{Name: "Seni, Budaya"},
		{Name: "Sosial, Hukum, Politik"},
		{Name: "Standup Comedy"},
		{Name: "Pendidikan, Beasiswa"},
		{Name: "Tech, Start-Up"},
		{Name: "Wisata & Liburan"},
		{Name: "Lainnya"},
	}

	for _, topic := range data {
		if err := db.Create(&topic).Error; err != nil {
			fmt.Printf("Error when create topic %s: %s\n", topic.Name, err)
		}
	}
}
