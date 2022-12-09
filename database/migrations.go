package database

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email     *string        `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username  string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Province struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Code      int            `gorm:"index:idx_code,unique;type:int(4);not null" json:"code"`
	Name      string         `gorm:"index:idx_name;type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type City struct {
	ID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProvinceID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"province_id"`
	Code       int            `gorm:"index:idx_code,unique;type:int(4);not null" json:"code"`
	Name       string         `gorm:"index:idx_name;type:varchar(255);not null" json:"name"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Province   Province       `gorm:"foreignKey:ProvinceID"`
}

type Job struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Sec struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Bottom    int            `gorm:"type:int(11);not null" json:"bottom"`
	Top       int            `gorm:"type:int(11);not null" json:"top"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Bank struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Code      int            `gorm:"index:idx_code;type:int(4);not null" json:"code"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type genderType string

const (
	MEN   genderType = "men"
	WOMEN genderType = "women"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}

func (gt genderType) Value() (driver.Value, error) {
	return string(gt), nil
}

type Member struct {
	ID            uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"user_id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	NationalityID string         `gorm:"type:varchar(255);not null" json:"nationality_id"`
	Gender        genderType     `gorm:"type:enum('men', 'women')" json:"gender"`
	Address       string         `gorm:"type:text;not null;" json:"address"`
	DateOfBirth   datatypes.Date `gorm:"type:date;not null;" json:"date_of_birth"`
	Qualifitacion string         `gorm:"type:varchar(255);not null" json:"qualification"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	User          User           `gorm:"foreignKey:UserID"`
}

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Province{})
	DB.AutoMigrate(&City{})
	DB.AutoMigrate(&Job{})
	DB.AutoMigrate(&Sec{})
	DB.AutoMigrate(&Bank{})
	DB.AutoMigrate(&Member{})
}
