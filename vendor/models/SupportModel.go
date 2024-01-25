package models

import "time"

type Support struct {
	SupportId 		int
	Area 			int
	Name 			string
	Email 			string
	Message 		string
	Open 			bool
	Created 		time.Time
}
