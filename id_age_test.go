package chinaid

import (
	"testing"
	"time"
)

func TestIDCardDetail_GetAge(t1 *testing.T) {
	type fields struct {
		Sex      int
		Birthday time.Time
		Addr     Addr
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "",
			fields: fields{},
			want:   31,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t,_ := IDCard("341302199006041233").Decode()

			if got := t.GetAge(); got != tt.want {
				t1.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}