package court_reserver_interface

import "time"

type Site int

const (
	XingqingSoutheastTennisCourts         Site = 301
	RollerCyclingCourts                   Site = 181
	TestCourts                            Site = 161
	XingqingShelteredTennisCourts         Site = 53
	XingqingCenter3rdFloorBadmintonCourts Site = 42
	RestrictedGym                         Site = 121
	XingqqingTableTennisCourts            Site = 43
	XingqingCenter1stFloorBadmintonCourts Site = 41
	XingqingCenterTennisCourts            Site = 55
	XingqingCenterSquashCourts            Site = 56
	IHarbourMainTennisCourts              Site = 82
	XingqingCenter1stFloorGym             Site = 44
	IHarbourMega3TableTennisCourts        Site = 105
	IHarbourMega2BadmintonCourts          Site = 103
	XingqingEastTennisCourts              Site = 52
	IHarbourMega3BadmintonCourts          Site = 104
	IHarbourMega1TableTennisCourts        Site = 102
	IHarbourMega1BadmintonCourts          Site = 101
	XingqingSouthTennisCourts             Site = 54
	MedicalTennisCourts                   Site = 51
	YantaFinancialTableTennisCourts       Site = 50
)

// The lookahead booking time of a site.
func SiteLookahead(site Site) int {
	switch site {
	case XingqingSoutheastTennisCourts:
		return 3
	case RollerCyclingCourts:
		return 5
	case TestCourts:
		return 7
	case XingqingShelteredTennisCourts:
		return 4
	case XingqingCenter3rdFloorBadmintonCourts:
		return 4
	case RestrictedGym:
		return 5
	case XingqqingTableTennisCourts:
		return 5
	case XingqingCenter1stFloorBadmintonCourts:
		return 4
	case XingqingCenterTennisCourts:
		return 4
	case XingqingCenterSquashCourts:
		return 4
	case IHarbourMainTennisCourts:
		return 2
	case XingqingCenter1stFloorGym:
		return 5
	case IHarbourMega3TableTennisCourts:
		return 3
	case IHarbourMega2BadmintonCourts:
		return 2
	case XingqingEastTennisCourts:
		return 2
	case IHarbourMega3BadmintonCourts:
		return 2
	case IHarbourMega1TableTennisCourts:
		return 3
	case IHarbourMega1BadmintonCourts:
		return 2
	case XingqingSouthTennisCourts:
		return 2
	case MedicalTennisCourts:
		return 2
	case YantaFinancialTableTennisCourts:
		return 5
	default:
		return 2
	}
}

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
	Failed
)

type ReservationStatus struct {
	Code      ReservationStatusCode
	Msg       string
	CourtTime map[string]string // time -> court
}

type CourtReserver interface {
	BookNow(time_zone *time.Location, reservation *Reservation, captcha_solver CaptchaSolver) ReservationStatus
}
