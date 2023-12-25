package auth

type Profile struct {
	ProfID   int    `db:"PROF_ID"`
	Login    string `db:"LOGIN"`
	Password string `db:"PASSWORD"`
	Token    string
}
