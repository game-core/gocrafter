package account

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
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

			err := handler(c)
			if tt.name != "正常：通過できる" {
				if !reflect.DeepEqual(err, tt.want(c)) {
					t.Errorf("ListExample() error = %v, wantErr %v", err, tt.want(c))
					return
				}
			} else {
				if !reflect.DeepEqual(http.StatusOK, rec.Code) {
					t.Errorf("ListExample() = %v, want %v", rec.Code, tt.want)
				}
			}
		})
	}
}
