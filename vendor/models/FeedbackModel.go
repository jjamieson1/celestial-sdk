package models

import "time"

type CustomerFeedback struct {
	FeedbackId		int			`json:"feedbackId"`
	OrderId			string		`json:"orderId"`
	Accurate		bool		`json:"accurate"`
	Complaint		string		`json:"complaint"`
	Confidence		int			`json:"confidence"`
	Satisfaction	int			`json:"satisfaction"`
	Feedback		string		`json:"feedback"`
	Created 		time.Time	`json:"created"`
}
