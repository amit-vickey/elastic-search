package elasticsearch

import "time"

type Student struct {
	RollNumber	 int64 	   `json:"roll_number"`
	Name         string    `json:"name"`
	Age          int64     `json:"age"`
	GPA			 float64   `json:"gpa"`
	JoinedOn	 time.Time `json:"joined_on"`
	IsActive	 bool 	   `json:"is_active"`
}
