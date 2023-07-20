package main

import (
	"database/sql"
	"testing"
)

func Test_calculatePopulation_empty(t *testing.T) {
	cities := []City{}
	got := calculatePopulation(cities)
	want := map[string]int64{}
	if len(want) != 0 {
		t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
	}
}

func Test_calculatePopulation_one(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
	}
	got := calculatePopulation(cities)
	want := map[string]int64{"JPN": 100}

	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
	}
	for c := range got {
		if got[c] != want[c] {
			t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
		}
	}
}

func Test_calculatePopulation_multi(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 400,
				Valid: true,
			},
		},
	}

	got := calculatePopulation(cities)
	want := map[string]int64{"JPN": 300, "USA": 400}

	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
	}
	for c := range got {
		if got[c] != want[c] {
			t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
		}
	}
}

func Test_calculatePopulation_invalid(t *testing.T) {
	cities := []City{
		{
			CountryCode: sql.NullString{
				String: "",
				Valid:  false,
			},
			Population: sql.NullInt64{
				Int64: 100,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "JPN",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 200,
				Valid: true,
			},
		},
		{
			CountryCode: sql.NullString{
				String: "USA",
				Valid:  true,
			},
			Population: sql.NullInt64{
				Int64: 400,
				Valid: true,
			},
		},
	}

	got := calculatePopulation(cities)
	want := map[string]int64{"JPN": 200, "USA": 400}

	if len(got) != len(want) {
		t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
	}
	for c := range got {
		if got[c] != want[c] {
			t.Errorf("calculatePopulation(%v) = %v, want %v.", cities, got, want)
		}
	}
}
