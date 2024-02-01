package user

import (
	clients2 "github.com/jjamieson1/celestial-sdk/clients"
)

func GetUserProfile(tenantId string, userId string) ([]byte, int, error) {
	url := "http://localhost:9100/api/profile/" + userId
	return clients2.CallRestEndPoint(url, "GET", tenantId, nil)
}

func GetUserProfileByUserName(tenantId string, userName string) ([]byte, int, error) {
	url := "http://localhost:9100/api/profile/username" + userName
	return clients2.CallRestEndPoint(url, "GET", tenantId, nil)
}

func GetAllUsers(tenantId string) ([]byte, int, error) {
	url := "http://localhost:9100/api/profile"

	return clients2.CallRestEndPoint(url, "GET", tenantId, nil)
}
