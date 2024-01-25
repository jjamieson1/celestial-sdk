package models

type AuditMessage struct {
	Type    string `json:"type"` // Log or Audit
	Service string `json:"service"`
	Subject string `json:"subject"` // Inject the subject
	Err     error  `json:"error"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Code    int    `json:"code"` // the error or audit may have a lookup code
}
