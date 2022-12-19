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

type Category struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"index:idx_name;type:varchar(255);not null" json:"name"`
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

type SecQuestion struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Question  string         `gorm:"type:text;not null;" json:"question"`
	IndexOf   uint           `gorm:"type:int(11) UNSIGNED;not null;default:0" json:"index_of"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type SecAnswer struct {
	ID            uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	SecQuestionID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"sec_question_id"`
	Answer        string         `gorm:"type:text;not null;" json:"question"`
	Score         int            `gorm:"type:int(11);not null;" json:"score"`
	IndexOf       uint           `gorm:"type:int(11) UNSIGNED;not null;default:0" json:"index_of"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	SecQuestion   SecQuestion    `gorm:"foreignKey:SecQuestionID"`
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

type Admin struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"user_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	IsActive  bool           `gorm:"type:tinyint(1);not null" json:"is_active"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Member struct {
	ID                  uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID              uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"user_id"`
	CityID              uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"city_id"`
	JobID               uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"job_id"`
	SecID               uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"sec_id"`
	BankID              uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;" json:"bank_id"`
	Name                string         `gorm:"type:varchar(255);not null" json:"name"`
	NationalityID       string         `gorm:"type:varchar(255);not null" json:"nationality_id"`
	Gender              genderType     `gorm:"type:enum('men', 'women')" json:"gender"`
	Address             string         `gorm:"type:text;not null;" json:"address"`
	DateOfBirth         datatypes.Date `gorm:"type:date;not null;" json:"date_of_birth"`
	Qualifitacion       string         `gorm:"type:varchar(255);not null" json:"qualification"`
	IsNatinalityIDValid bool           `gorm:"type:tinyint(1);not null" json:"is_nationality_id_valid"`
	IsActive            bool           `gorm:"type:tinyint(1);not null" json:"is_active"`
	IsSuspend           bool           `gorm:"type:int(11);not null" json:"is_suspend"`
	Point               int            `gorm:"type:tinyint(1);not null" json:"point"`
	AccountNumber       string         `gorm:"type:varchar(255);not null" json:"account_number"`
	AccountHolder       string         `gorm:"type:varchar(255);not null" json:"account_holder"`
	PhoneModel          string         `gorm:"type:varchar(255);not null" json:"phone_model"`
	DeviceToken         string         `gorm:"type:varchar(255);not null" json:"device_token"`
	OtpCode             string         `gorm:"type:varchar(255);not null" json:"otp_code"`
	OtpExpiredAt        time.Time      `gorm:"column:otp_expired_at;null" json:"otp_expired_at"`
	LastActive          time.Time      `gorm:"column:last_active;null" json:"last_active"`
	CreatedAt           time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	User                User           `gorm:"foreignKey:UserID"`
	City                City           `gorm:"foreignKey:CityID"`
	Job                 Job            `gorm:"foreignKey:JobID"`
	Sec                 Sec            `gorm:"foreignKey:SecID"`
	Bank                Bank           `gorm:"foreignKey:BankID"`
}

type Campaign struct {
	ID               uuid.UUID       `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Title            string          `gorm:"type:varchar(255);not null" json:"title"`
	Description      string          `gorm:"type:text;not null" json:"description"`
	ShortDescription string          `gorm:"type:text;not null" json:"short_description"`
	Image            *string         `gorm:"type:text;" json:"image"`
	StartAt          *datatypes.Date `gorm:"column:start_at;type:date;" json:"start_at"`
	FinishAt         *datatypes.Date `gorm:"column:finish_at;type:date;" json:"finish_at"`
	Status           uint            `gorm:"column:status;type:smallint(6);not null;" json:"status"`
	Background       string          `gorm:"type:varchar(255);not null" json:"background"`
	CreatedAt        time.Time       `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt  `gorm:"column:deleted_at;" json:"deleted_at"`
}

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Province{})
	DB.AutoMigrate(&City{})
	DB.AutoMigrate(&Job{})
	DB.AutoMigrate(&Sec{})
	DB.AutoMigrate(&SecQuestion{})
	DB.AutoMigrate(&SecAnswer{})
	DB.AutoMigrate(&Bank{})
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Member{})
	DB.AutoMigrate(&Campaign{})
	DB.Exec("ALTER TABLE `campaigns` CHANGE `start_at` `start_at` DATE NULL;")
	DB.Exec("ALTER TABLE `campaigns` CHANGE `finish_at` `finish_at` DATE NULL;")
}
