package vec

import "testing"

func TestMag(t *testing.T) {
	cases := []struct {
		in  Vec2
		out float64
	}{
		{Vec2{3, 4}, 5},
		{Vec2{0.6, 0.8}, 1},
		{Vec2{}, 0},
	}

	for _, c := range cases {
		got := c.in.Mag()
		if got != c.out {
			t.Errorf("Mag(%v) == %v, wanted %v", c.in, got, c.out)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, want Vec2
	}{
		{Vec2{0, 0}, Vec2{1, 1}, Vec2{1, 1}},
		{Vec2{3, 4}, Vec2{7, 8}, Vec2{10, 12}},
		{Vec2{}, Vec2{}, Vec2{}},
	}

	for _, c := range cases {
		got := c.a.Add(c.b)
		if got != c.want {
			t.Errorf("%v.Add(%v) == %v, wanted %v", c.a, c.b, got, c.want)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		a, b, want Vec2
	}{
		{Vec2{0, 0}, Vec2{1, 1}, Vec2{-1, -1}},
		{Vec2{3, 4}, Vec2{7, 8}, Vec2{-4, -4}},
		{Vec2{}, Vec2{}, Vec2{}},
	}

	for _, c := range cases {
		got := c.a.Sub(c.b)
		if got != c.want {
			t.Errorf("%v.Sub(%v) == %v, wanted %v", c.a, c.b, got, c.want)
		}
	}
}
