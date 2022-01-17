package numutil

import "testing"

func TestIntIf(t *testing.T) {
	cases := []struct {
		Name           string
		Condition      bool
		A, B, Expected int
	}{
		{"true", true, 1, 2, 1},
		{"false", false, 3, 4, 4},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := IntIf(c.Condition, c.A, c.B); ans != c.Expected {
				t.Fatalf("IntIf(%v, %d, %d) expected %d, but get %d",
					c.Condition, c.A, c.B, c.Expected, ans)
			}
		})
	}
}

func TestIntLeast(t *testing.T) {
	cases := []struct {
		Name               string
		N, Lower, Expected int
	}{
		{"9 least 10", 9, 10, 10},
		{"10 least 10", 10, 10, 10},
		{"11 least 10", 11, 10, 11},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := IntLeast(c.N, c.Lower); ans != c.Expected {
				t.Fatalf("IntLeast(%d, %d) expected %d, but get %d",
					c.N, c.Lower, c.Expected, ans)
			}
		})
	}
}

func TestIntMost(t *testing.T) {
	cases := []struct {
		Name               string
		N, Upper, Expected int
	}{
		{"9 most 10", 9, 10, 9},
		{"10 most 10", 10, 10, 10},
		{"11 most 10", 11, 10, 10},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := IntMost(c.N, c.Upper); ans != c.Expected {
				t.Fatalf("IntMost(%d, %d) expected %d, but get %d",
					c.N, c.Upper, c.Expected, ans)
			}
		})
	}
}

func TestIntBetween(t *testing.T) {
	cases := []struct {
		Name                      string
		N, Lower, Upper, Expected int
	}{
		{"9 between [10, 20]", 9, 10, 20, 10},
		{"10 between [10, 20]", 10, 10, 20, 10},
		{"15 between [10, 20]", 15, 10, 20, 15},
		{"20 between [10, 20]", 20, 10, 20, 20},
		{"21 between [10, 20]", 21, 10, 20, 20},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := IntBetween(c.N, c.Lower, c.Upper); ans != c.Expected {
				t.Fatalf("IntBetween(%d, %d, %d) expected %d, but get %d",
					c.N, c.Lower, c.Upper, c.Expected, ans)
			}
		})
	}
}
