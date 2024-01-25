package utilities

import (
	"net/http"
	"os"

	"github.com/jjamieson1/eden-sdk/models"
)

func GetReleaseInfo(w http.ResponseWriter, r *http.Request) {
	logger := Logger(r.Context())
	defer logger.Sync()
	var releaseInformation = models.ReleaseInformation{
		VersionNumber: os.Getenv("VERSION_NUMBER"),
		AppName:       os.Getenv("APP_NAME"),
		TenantName:    os.Getenv("TENANT_NAME"),
		StartedOn:     os.Getenv("STARTED_ON"),
	}
	WriteJSON(releaseInformation, http.StatusOK, w)
}
