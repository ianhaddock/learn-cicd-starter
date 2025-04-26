package auth 

import (
    "testing"
    "fmt"
    "net/http"
    "errors"
)


func TestGetAPIKey(t *testing.T) {

    cases := []struct {
        key string
        value string
        expected string
        expectedErr error
    }{
        {
            key: "Authorization",
            value: "ApiKey asfasfasfd",
            expected: "asfasfasfd",
            expectedErr: nil,
        },
        {
            expected: "",
            expectedErr: ErrNoAuthHeaderIncluded,
        },
        {
            key: "Authorization",
            value: "-",
            expected: "",
            expectedErr: errors.New("malformed authorization header"),
        },
        {
            key: "Authorization",
            value: "AiKy asfasfasfd",
            expected: "",
            expectedErr: errors.New("malformed authorization header"),
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test GetAPIKey Case #%v:", i), func(t *testing.T) {
            header := http.Header{}
            header.Add(c.key, c.value)

            actual, err := GetAPIKey(header)

            if err != nil {
                if err.Error() != c.expectedErr.Error() {
                    t.Errorf("Got unexpected error: '%v' expected: '%v'", err, c.expectedErr)
                    return
                }
            }

            if actual != c.expected {
                t.Errorf("expected %v does not match actual: %v", c.expected, actual)
                return
            }
        })
    }
}
