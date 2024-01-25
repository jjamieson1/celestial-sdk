package notification

import (
	"bytes"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
)

func SendViaSMTP(to string, toName string, subject string, templateVariables map[string]string, templateId int) error {
	from := "info@canwestcheer.ca"
	pass := "goFIGHTwin2468!"
	msg := createBodyFromTemplate(templateVariables, templateId, from, to, subject)
	err := smtp.SendMail("smtp.ionos.com:587",
		smtp.PlainAuth("", from, pass, "smtp.ionos.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return err
}

func createBodyFromTemplate(templateVariables map[string]string, templateId int, from, to, subject string) string {

	var bodyMessage string

	if templateId == 1 {
		log.Printf("Template Support response selected")
		bodyMessage = bodyMessage + `<p><img width="200px" src="https://canwestcheer.ca/wp-content/uploads/2020/10/Canwest-Virtual-Logo-Small.png" /></p>`
		bodyMessage = bodyMessage + "<strong>Dear " + templateVariables["name"] + "</strong>"
		bodyMessage = bodyMessage + "<br />"
		bodyMessage = bodyMessage + templateVariables["message"]
		bodyMessage = bodyMessage + "<hr />"
	} else if templateId == 2 {
		log.Printf("Template Organization Registration response selected")

		bodyMessage = bodyMessage + `<table width="100%" border="0" cellspacing="0" cellpadding="0"><tr><td align="center"><img width="200px" src="https://canwestcheer.ca/wp-content/uploads/2020/10/Canwest-Virtual-Logo-Small.png" /></td></tr></table>`
		bodyMessage = bodyMessage + "<h5>Dear " + templateVariables["firstName"] + "</h5>"
		bodyMessage = bodyMessage + "<h3>Thank you for your registration</h3>"
		bodyMessage = bodyMessage + "<p>Your registration has been received and we are getting ready to review and approve it.  Once this is done, we will send you some login information so you can come back and manage your team.</p><hr />"
		bodyMessage = bodyMessage + "<p>Your gym/school registered is: <strong>" + templateVariables["name"] + "</strong></p>"
		bodyMessage = bodyMessage + "<hr />"
	} else if templateId == 3 {
		log.Printf("Admin for organization response selected")
		bodyMessage = bodyMessage + `<table width="100%" border="0" cellspacing="0" cellpadding="0"><tr><td align="center"><img width="200px" src="https://canwestcheer.ca/wp-content/uploads/2020/10/Canwest-Virtual-Logo-Small.png" /></td></tr></table>`
		bodyMessage = bodyMessage + "<strong>Dear " + templateVariables["name"] + "</strong>"
		bodyMessage = bodyMessage + "<h3>You account is approved and ready</h3>"
		bodyMessage = bodyMessage + "<p>" + templateVariables["message"] + "</p>"
		bodyMessage = bodyMessage + "<p>Now you can login, and manage your teams, invite other users to do tasks on your behalf and register for events.  Please come back to the site and select login to request a magic link to login. </p><hr />"
		bodyMessage = bodyMessage + "<hr />"
		bodyMessage = bodyMessage + "Your organization: <h3>" + templateVariables["commonName"] + "</h3>"
	} else if templateId == 4 {
		log.Printf("Template feedback response selected")
		bodyMessage = bodyMessage + `<p><img width="200px" src="https://canwestcheer.ca/wp-content/uploads/2020/10/Canwest-Virtual-Logo-Small.png" /></p>`
		bodyMessage = bodyMessage + "<strong>Dear " + templateVariables["name"] + "</strong>"
		bodyMessage = bodyMessage + "<br /><br />"
		bodyMessage = bodyMessage + "<p>Eden really appreciates you as a customer.  We are committed to providing the highest quality products and service</p><hr />"
		bodyMessage = bodyMessage + "<p>If you could spend a minute and let us know how we did on processing your last order, We would really appreciate it, it will help us improve your service for the future.</p>"
		bodyMessage = bodyMessage + "<hr />"
		bodyMessage = bodyMessage + "Please follow the link below: <a href=\"https://store.eden.green/App/Feedback/" + templateVariables["code"] + "\">Eden Store Feedback form</a>"
	} else if templateId == 5 {
		log.Printf("Template email login response selected")
		bodyMessage = bodyMessage + `<p><img width="200px" src="https://canwestcheer.ca/wp-content/uploads/2020/10/Canwest-Virtual-Logo-Small.png" /></p>`
		bodyMessage = bodyMessage + "<strong>Greetings</strong>"
		bodyMessage = bodyMessage + "<br /><br />"
		bodyMessage = bodyMessage + "<p>If you just recently requested a login link, the link to log you in is below.</p><hr />"
		bodyMessage = bodyMessage + "<p>If you did not request a token, please disregard.  This token will be invalid in 5 minutes.</p>"
		bodyMessage = bodyMessage + "<hr />"
		bodyMessage = bodyMessage + "Please click the following to login: <a href=\"https://comp.sca.ca/en/app/login/email/" + to + "/token/" + templateVariables["token"] + "\">Your Magic Link</a>"
	}

	//"code": code,
	//	"email": email,s
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message

}
