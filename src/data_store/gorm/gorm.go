package gorm

//The fantastic ORM library for Golang, aims to be developer friendly

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/utility/error"

	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"           // postgres
	_ "github.com/mattn/go-sqlite3" // sqllite3
)

// Gorm is struct for singleton
type Gorm struct {
	DB *gorm.DB // Session
}

// Global instance
var Instance = &Gorm{DB: nil}

func InitGorm() *gorm.DB {

	db, err := gorm.Open(config.RDB_TYPE, "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		log.panic(error.ErrNotFountInstant)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	Instance.DB = db

	return db
}

func Connector(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		if Instance.Session != nil {
			log.panic(error.ErrNotFountInstant)
		}

		s := Instance.Session.Clone()
		defer s.Close()
		c.Set("mongodb", s)

		next(c)
	})
}

// Example struct
// // Define Models (Structs)

// type User struct {
// 	ID        int
// 	Birthday  time.Time
// 	Age       int
// 	Name      string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
// 	Num       int    `sql:"AUTO_INCREMENT"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time

// 	Emails            []Email       // One-To-Many relationship (has many)
// 	BillingAddress    Address       // One-To-One relationship (has one)
// 	BillingAddressID  sql.NullInt64 // Foreign key of BillingAddress
// 	ShippingAddress   Address       // One-To-One relationship (has one)
// 	ShippingAddressID int           // Foreign key of ShippingAddress
// 	IgnoreMe          int           `sql:"-"`                          // Ignore this field
// 	Languages         []Language    `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
// }

// type Email struct {
// 	ID         int
// 	UserID     int    `sql:"index"`                          // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
// 	Email      string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
// 	Subscribed bool
// }

// type Address struct {
// 	ID       int
// 	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
// 	Address2 string         `sql:"type:varchar(100);unique"`
// 	Post     sql.NullString `sql:"not null"`
// }

// type Language struct {
// 	ID   int
// 	Name string `sql:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
// 	Code string `sql:"index:idx_name_code"` // `unique_index` also works
// }
