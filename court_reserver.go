package court_reserver_interface

import (
	"time"

	"github.com/endaytrer/court_reserver_interface/captcha_solver"
)

type Site int

type SingleBook struct {
	// book all contiguous courts COVERINGs Date + StartTime to Date + StartTime + Duration. Time starts with UTC.
	StartTime time.Duration
	Duration  time.Duration
	// preferring booking name
	CourtNamePreference []string
}
type Reservation struct {
	Date time.Time
	Site Site
	// book ONLY the top book available
	Preferences []SingleBook
	Priority    int // lower priority value would try to book first.
}

type ReservationStatusCode int

const (
	Pending ReservationStatusCode = iota
	Success
	NeedPayment
	Failed
	NeedAuthorization
)

type ReservationStatus struct {
	Code      ReservationStatusCode
	Msg       string
	CourtTime map[string]string // time -> court
}

type CourtReserver interface {
	BookNow(time_zone *time.Location, reservation *Reservation, captcha_solver captcha_solver.CaptchaSolver, payment_password *string) ReservationStatus
}
