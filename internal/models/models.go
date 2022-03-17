package models

import "time"

//Users is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Rooms is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Restrictions is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Reservations is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
}

//RoomRestriction is the restriction model
type RoomRestriction struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	Room          Room
	ReservationID int
	Reservation   Reservation
	RestrictionID int
	Restriction   Restriction
}
