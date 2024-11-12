package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age      int
	Favorite string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Henry", 32},
			[]string{"Henry"},
		},
		{
			"nested struct",
			Person{
				"Henry",
				Profile{32, "Fitness"},
			},
			[]string{"Henry", "Fitness"},
		},
		{
			"pointers to things",
			&Person{
				"Henry",
				Profile{32, "Fitness"},
			},
			[]string{"Henry", "Fitness"},
		},
		{
			"slices",
			[]Profile{
				{32, "Fitness"},
				{18, "Bakery"},
			},
			[]string{"Fitness", "Bakery"},
		},
		{
			"array",
			[2]Profile{
				{32, "Fitness"},
				{18, "Bakery"},
			},
			[]string{"Fitness", "Bakery"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	assertContains := func(t testing.TB, got []string, expected string) {
		t.Helper()
		contains := slices.Contains(got, expected)
		if !contains {
			t.Errorf("expected %v to contain %q but it didn't", got, expected)
		}
	}

	t.Run("with maps", func(t *testing.T) {
		mapInstance := map[string]string{
			"LastName":   "Henry",
			"FamilyName": "Chou",
		}

		var got []string
		Walk(mapInstance, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Henry")
		assertContains(t, got, "Chou")
	})

	t.Run("with channels", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{33, "Fitness"}
			channel <- Profile{18, "Reading"}
			close(channel)
		}()

		var got []string
		want := []string{"Fitness", "Reading"}

		Walk(channel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		myFunction := func() (Profile, Profile) {
			return Profile{33, "Fitness"}, Profile{18, "Reading"}
		}

		var got []string
		want := []string{"Fitness", "Reading"}

		Walk(myFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
