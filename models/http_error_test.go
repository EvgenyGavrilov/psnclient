package models

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHTTPError_Error(t *testing.T) {
	statusCode := 500
	err := &HTTPError{StatusCode: statusCode}
	require.EqualError(t, err, fmt.Sprintf("Response http status %d", statusCode))
}
