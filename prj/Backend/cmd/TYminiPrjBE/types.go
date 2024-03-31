package main

// -------------------------------Types----------------------------
type loginDetails struct {
	username string
	password string
}

type userDetails struct {
	Username string
	Id       uint32
	TicketId uint64
	UserQR   string
}

type recoveryDetails struct {
	ReqSent   bool
	Email     string
	Auth      string
	AuthCode  uint32
	MsgString string
}

type registerDbDetails struct {
	id       uint32
	username string
	email    string
	passHsh  string
	salt     string
	ticketId uint64
}

type registerDetails struct {
	email    string
	username string
	password string
}

type bookingFormIp struct {
	tier1 uint32
	tier2 uint32
	tier3 uint32
}

type bookingTicketFeedback struct {
	AddedToCart  bool
	Checkout     bool
	MsgString    string
	Premium      uint32
	Base         uint32
	Minimum      uint32
	PremiumCost  uint32
	BaseCost     uint32
	MinimumCost  uint32
	PremiumTotal uint32
	BaseTotal    uint32
	MinimumTotal uint32
}

// -------------------------------Variables----------------------------
var user = userDetails{}
var userAuth = recoveryDetails{}
var email string
var apass string
var port string

// -------------------------------Constants-------------------------------
const t1Cost uint32 = 500
const t2Cost uint32 = 300
const t3Cost uint32 = 100
