package Model

type Employee struct {
	ID            uint   `json: "ID"`
	FullName      string `json: "fullname"`
	Male          bool   `json: "male"`
	Nationality   string `json: "nationality"`
	DirectManager uint   `json: "directmanager"`
	Password      string `json: "pass"`
}
