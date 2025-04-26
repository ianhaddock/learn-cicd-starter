package main

import (
    "testing"
    "net/http"
    "errors"
    "github.com/bootdotdev/learn-cicd-starter/internal/auth"
)


func TestGetAPIKey(t *testing.T) {

    cases := []struct {
        input http.Header
        expected string
        expectedErr error
    }{
        {
            input: http.Header{},
            expected: "",
            expectedErr: auth.ErrNoAuthHeaderIncluded,
        },
        {
            input: http.Header{
                "Authorization": []string{"ApiKey asfasfasfd"},
            },
            expected: "asfasfasfd",
            expectedErr: nil,
        },
        {
            input: http.Header{
                "Authorization": []string{"AiKy asfasfasfd"},
            },
            expected: "",
            expectedErr: errors.New("malformed authorization header"),
        },
    }

    for _, c := range cases {
        actual, err := auth.GetAPIKey(c.input) //headers)
        if err != nil {
            if err.Error() != c.expectedErr.Error() {
                t.Errorf("Got unexpected error: '%v' expected: '%v'", err, c.expectedErr)
            }
        }
        if len(actual) != len(c.expected) {
            t.Errorf("expected length %d does not match actual length %d", len(actual), len(c.expected))
        }
    }
}
