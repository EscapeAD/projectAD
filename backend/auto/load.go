package auto

import (
	"escapead/projectAD/backend/api/database"
	"escapead/projectAD/backend/api/models"
	"escapead/projectAD/backend/api/utils/console"
	"log"
)

// Load - Load data
func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Posts
	err = db.Debug().DropTableIfExists(&models.Post{}).Error
	if err != nil {
		log.Fatal(err)
	}
	// Users
	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.Post{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
	for i, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i].Author).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(user)
	}
}
