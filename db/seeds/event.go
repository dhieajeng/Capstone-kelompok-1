package seeds

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"gorm.io/gorm"
	"time"
)

func CreateEventSeeder(db *gorm.DB) {
	user := new(entity.User)
	location := new(entity.Location)
	category := new(entity.Category)
	topic := new(entity.Topic)
	if err := db.First(location).Error; err != nil {
		fmt.Printf("Error finding location: %s", err)
	}
	if err := db.First(category).Error; err != nil {
		fmt.Printf("Error finding category: %s", err)
	}
	if err := db.First(topic).Error; err != nil {
		fmt.Printf("Error finding topic: %s", err)
	}
	if err := db.First(user).Error; err != nil {
		fmt.Printf("Error finding user: %s", err)
	}

	data := []entity.NewEventParams{
		{
			Name:             "Saloka Fest 2024",
			Start:            time.Date(2024, 06, 22, 10, 0, 0, 0, time.Local),
			End:              time.Date(2024, 06, 22, 12, 0, 0, 0, time.Local),
			Address:          "Saloka Theme Park, Jawa Tengah",
			AddressLink:      "https://www.google.com/maps/search/?api=1&query=-7.28074,110.459",
			Organizer:        "PT Panorama Indah Permai",
			Description:      "Saloka Fest di tahun 2024 ini sangat spesial karena bertepatan dengan perayaan ulang tahun Saloka Theme Park yang ke-5. Tema MUSIC & ART akan menjadi anchor point untuk event kedepannya dan sustainability menjadikampanye yang akan dikomunikasikan melalui event Saloka Fest.",
			TermAndCondition: "Sing penting bayar",
			IsPaid:           true,
			IsPublic:         true,
			IsApproved:       true,
			UserID:           user.ID,
			LocationID:       location.ID,
			CategoryID:       category.ID,
			TopicID:          topic.ID,
		},
		{
			Name:             "WONDERFUL UNGASAN",
			Start:            time.Date(2024, 06, 29, 17, 0, 0, 0, time.Local),
			End:              time.Date(2024, 06, 29, 23, 55, 0, 0, time.Local),
			Address:          "Lapangan Bola Bantang",
			AddressLink:      "https://www.google.com/maps/search/?api=1&query=-8.84035,115.164",
			Organizer:        "Wonderful Ungasan",
			Description:      "Wonderful Ungasan merupakan wadah yang merepresentasikan generasi dengan pelestarian kearifan lokal yang dikemas dengan tampilan informasi tentang entertainment lebih menarik. Event Wonderful Ungasan ini terbentuk sejak 2018.",
			TermAndCondition: "Sing penting bayar",
			IsPaid:           true,
			IsPublic:         true,
			IsApproved:       true,
			UserID:           user.ID,
			LocationID:       location.ID,
			CategoryID:       category.ID,
			TopicID:          topic.ID,
		},
	}

	for _, event := range data {
		event := entity.NewEvent(event)
		if err := db.Create(&event).Error; err != nil {
			fmt.Printf("Error when create event %s: %s\n", event.Name, err)
		}
	}
}
