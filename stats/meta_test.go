package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMeta(t *testing.T) {
	tests := []struct {
		name    string
		want    *Meta
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMeta()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
