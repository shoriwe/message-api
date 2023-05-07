package http_utils

import (
	"encoding/base64"
	"fmt"
)

func ToBearer(token string) string {
	return fmt.Sprintf("Bearer %s", base64.StdEncoding.EncodeToString([]byte(token)))
}
