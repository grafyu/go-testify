package cafe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func servAns(count, city string) string {
	reqURL := fmt.Sprintf("/cafe?count=%s&city=%s", count, city)

	req := httptest.NewRequest("GET", reqURL, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return fmt.Sprintf("%d %s", responseRecorder.Code, responseRecorder.Body.String())
}

func Test_mainHandle(t *testing.T) {
	totalCount := 4
	type args struct {
		count string
		city  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// TODO: Add test cases.
			name: "correct city, not null count: code 200, send is not empty",
			args: args{count: "3", city: "moscow"},
			want: "200 Мир кофе,Сладкоежка,Кофе и завтраки",
		},
		{
			// TODO: Add test cases.
			name: "wrong city, not null count: code 400, send - wrong city value",
			args: args{count: "3", city: "leningrad"},
			want: "400 wrong city value",
		},
		{
			// TODO: Add test cases.
			name: "correct city, over the maximum count: code 200, send all cafes",
			args: args{count: strconv.Itoa(totalCount + 1), city: "moscow"},
			want: "200 Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
		},
		{
			name: "correct city, empty count: code 400, send - count missing",
			args: args{count: "", city: "moscow"},
			want: "400 count missing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, servAns(tt.args.count, tt.args.city),
				"The response to %s and %s", tt.args.count, tt.args.city)
		})
	}
}
