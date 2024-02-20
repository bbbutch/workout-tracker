package database

import (
	"errors"
	"regexp"

	"github.com/jovandeginste/workout-tracker/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUsernameInvalidLength = errors.New("username has invalid length")
	ErrUsernameInvalidChars  = errors.New("username contains invalid characters")

	UsernameValidRegex = regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
)

type User struct {
	gorm.Model
	Password string `form:"-" gorm:"type:varchar(128);not null"`
	Salt     string `form:"-" gorm:"type:varchar(16);not null"`
	Username string `form:"username" gorm:"uniqueIndex;not null;type:varchar(20)"`
	Name     string `form:"name" gorm:"type:varchar(64);not null"`
	Active   bool   `form:"active"`
	Admin    bool   `form:"admin"`

	Profile  Profile
	Workouts []Workout
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, db.Error
	}

	return u, nil
}

func GetUserByID(db *gorm.DB, userID int) (*User, error) {
	var u User

	if err := db.First(&u, userID).Error; err != nil {
		return nil, db.Error
	}

	return &u, nil
}

func GetUser(db *gorm.DB, username string) (*User, error) {
	var u User

	if err := db.Where(&User{Username: username}).First(&u).Error; err != nil {
		return nil, db.Error
	}

	return &u, nil
}

type Profile struct {
	gorm.Model
	UserID int
	Theme  ThemePreference
}

type ThemePreference string

func (u *User) IsActive() bool {
	if u == nil {
		return false
	}

	if !u.Active || u.Password == "" || u.Username == "" {
		return false
	}

	return true
}

func (u *User) ValidLogin(password string) bool {
	if !u.IsActive() {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.AddSalt(password))) == nil
}

func (u *User) AddSalt(password string) string {
	return u.Salt + password
}

func (u *User) IsValid() error {
	if len(u.Username) < 3 || len(u.Username) > 20 {
		return ErrUsernameInvalidLength
	}

	if !UsernameValidRegex.MatchString(u.Username) {
		return ErrUsernameInvalidChars
	}

	return nil
}

func (u *User) SetPassword(password string) error {
	if err := u.GenerateSalt(); err != nil {
		return err
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.AddSalt(password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(cryptedPassword)

	return nil
}

func (u *User) Create(db *gorm.DB) error {
	if err := u.GenerateSalt(); err != nil {
		return err
	}

	return db.Create(u).Error
}

func (u *User) GenerateSalt() error {
	if u.Salt != "" {
		return nil
	}

	var err error

	if u.Salt, err = util.GenerateRandomString(8); err != nil {
		return err
	}

	return nil
}

func (u *User) Save(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(u).Error
}

func (u *User) GetWorkout(db *gorm.DB, id int) (*Workout, error) {
	var w *Workout

	if err := db.Where(&Workout{UserID: u.ID}).First(&w, id).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) GetWorkouts(db *gorm.DB) ([]Workout, error) {
	var w []Workout

	if err := db.Where(&Workout{UserID: u.ID}).Order("date DESC").Find(&w).Error; err != nil {
		return nil, err
	}

	return w, nil
}

func (u *User) AddWorkout(db *gorm.DB, workoutType, notes string, content []byte) (*Workout, error) {
	w := NewWorkout(u, workoutType, notes, content)

	if err := w.Create(db); err != nil {
		return nil, err
	}

	return w, nil
}
