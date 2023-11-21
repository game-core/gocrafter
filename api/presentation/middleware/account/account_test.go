package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountMiddleware_AccountMiddleware(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		{
			name: "正常：通過できる",
			args: args{
				next: func(c echo.Context) error {
					return c.String(http.StatusOK, "This should be reached")
				},
			},
			want: func(c echo.Context) error {
				return c.String(http.StatusOK, "This should be reached")
			},
		},
		{
			name: "異常：エラー（Authorization header is missing）",
			args: args{
				next: func(c echo.Context) error {
					return c.String(http.StatusOK, "This should not be reached")
				},
			},
			want: func(c echo.Context) error {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is missing")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccountMiddleware()
			handler := a.AccountMiddleware(tt.args.next)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			assert.NotNil(t, tt.want, "Expected handler function is nil")

			err := handler(c)
			if tt.name == "正常：通過できる" {
				assert.Equal(t, http.StatusOK, rec.Code)
			} else {
				assert.EqualError(t, err, tt.want(c).(*echo.HTTPError).Error())
			}
		})
	}
}
