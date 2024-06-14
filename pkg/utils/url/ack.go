package url

import "fmt"

const UserAccuserAccessionKey = "USER_ACCUSER_ACCESSION_KEY"

func GetAccessToken(id int64) string {
	return fmt.Sprintf("%s%d", UserAccuserAccessionKey, id)
}
