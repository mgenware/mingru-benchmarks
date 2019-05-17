package gormExample

import "time"

type Employee struct {
	EmpNo     int       `gorm:"primary_key"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Gender    string    `gorm:"not null"`
	HireDate  time.Time `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
}

type Title struct {
	Title    string    `gorm:"not null"`
	EmpNo    int       `gorm:"not null"`
	Employee Employee  `gorm:"ForeignKey:EmpNo;AssociationForeignKey:EmpNo"`
	FromDate time.Time `gorm:"not null"`
	ToDate   time.Time `gorm:"not null"`
}
