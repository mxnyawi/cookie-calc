package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCookies(t *testing.T) {
	tests := []struct {
		name    string
		data    [][]string
		want    []*Cookie
		wantErr bool
	}{
		{
			name: "valid data",
			data: [][]string{
				{"AtY0laUfhglK3lC7", "2018-12-09T14:19:00+00:00"},
				{"SAZuXPGUrfbcn5UA", "2018-12-09T10:13:00+00:00"},
			},
			want: []*Cookie{
				{Value: "AtY0laUfhglK3lC7", Date: "2018-12-09T14:19:00+00:00"},
				{Value: "SAZuXPGUrfbcn5UA", Date: "2018-12-09T10:13:00+00:00"},
			},
			wantErr: false,
		},
		{
			name: "invalid cookie value",
			data: [][]string{
				{"AtY0laUfhglK3lC7", "2018-12-09T14:19:00+00:00"},
				{"invalid", "2018-12-09T10:13:00+00:00"},
			},
			want: []*Cookie{
				{Value: "AtY0laUfhglK3lC7", Date: "2018-12-09T14:19:00+00:00"},
				nil,
			},
			wantErr: false,
		},
		{
			name: "missing date",
			data: [][]string{
				{"AtY0laUfhglK3lC7", "2018-12-09T14:19:00+00:00"},
				{"SAZuXPGUrfbcn5UA"},
			},
			want: []*Cookie{
				{Value: "AtY0laUfhglK3lC7", Date: "2018-12-09T14:19:00+00:00"},
				nil,
			},
			wantErr: false,
		},
		{
			name:    "empty data",
			data:    [][]string{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractCookies(tt.data)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestExtractDay(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
	}{
		{
			name:     "positive: correctly formatted date",
			date:     "2018-12-09T10:13:00+00:00",
			expected: "2018-12-09",
		},
		{
			name:     "negative: incorrectly formatted date",
			date:     "2018-12-09",
			expected: "2018-12-09",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractDay(tt.date)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		cookies [][]string
		date    string
		want    []string
		wantErr error
	}{
		{
			name: "Single cookie on given date",
			cookies: [][]string{
				{"SAZuXPGUrfbcn5UA", "2022-01-01T00:00:00Z"},
			},
			date: "2022-01-01",
			want: []string{"SAZuXPGUrfbcn5UA"},
		},
		{
			name: "Multiple cookies on given date",
			cookies: [][]string{
				{"SAZuXPGUrfbcn5UA", "2022-01-01T00:00:00Z"},
				{"5UAVanZf6UtGyKVS", "2022-01-01T00:00:00Z"},
				{"SAZuXPGUrfbcn5UA", "2022-01-01T00:00:00Z"},
			},
			date: "2022-01-01",
			want: []string{"SAZuXPGUrfbcn5UA"},
		},
		{
			name: "No cookies on given date",
			cookies: [][]string{
				{"SAZuXPGUrfbcn5UA", "2022-01-02T00:00:00Z"},
			},
			date:    "2022-01-01",
			wantErr: ErrNoCookies,
		},
		{
			name: "Invalid data",
			cookies: [][]string{
				{"cookies", "2022-01-01T00:00:00Z"},
			},
			date:    "2022-01-01",
			wantErr: ErrInvalidData,
		},
		{
			name: "Multiple cookies with same count on given date",
			cookies: [][]string{
				{"SAZuXPGUrfbcn5UA", "2022-01-01T00:00:00Z"},
				{"5UAVanZf6UtGyKVS", "2022-01-01T00:00:00Z"},
			},
			date: "2022-01-01",
			want: []string{"SAZuXPGUrfbcn5UA", "5UAVanZf6UtGyKVS"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.cookies, tt.date)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
