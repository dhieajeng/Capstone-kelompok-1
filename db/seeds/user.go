package seeds

import (
	"fmt"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserSeeds(db *gorm.DB) {
	getStringPointer := func(s string) *string {
		return &s
	}

	male := entity.Male
	female := entity.Female

	data := map[string]entity.User{
		"admin": {
			ID:       uuid.New(),
			Name:     "Gue Admin",
			Email:    "admin@mail.com",
			Password: "password",
			Role:     entity.Admin,
			Phone:    getStringPointer(faker.Phonenumber()),
			Address:  getStringPointer(faker.DomainName()),
			Avatar:   getStringPointer(faker.Sentence()),
			Birthdate: func() *time.Time {
				date, _ := time.Parse("2006-01-02", faker.Date())
				return &date
			}(),
			Gender: &male,
		},
		"buyer": {
			ID:       uuid.New(),
			Name:     "Gue Buyer",
			Email:    "buyer@mail.com",
			Password: "password",
			Role:     entity.Buyer,
			Phone:    getStringPointer(faker.Phonenumber()),
			Address:  getStringPointer(faker.DomainName()),
			Avatar:   getStringPointer(faker.Sentence()),
			Birthdate: func() *time.Time {
				date, _ := time.Parse("2006-01-02", faker.Date())
				return &date
			}(),
			Gender: &female,
		},
	}

	for _, user := range data {
		pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("Error when create user %s: %s\n", user.Name, err)
		}
		data := entity.NewUser(user.Name,
			user.Email,
			string(pw),
			user.Phone,
			user.Address,
			user.Avatar,
			user.Birthdate,
			user.Gender,
			user.Role,
		)

		if err := db.Create(&data).Error; err != nil {
			fmt.Printf("Error when create user %s: %s\n", data.Name, err)
			return
		}
	}
}
