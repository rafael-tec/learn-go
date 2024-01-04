package standard

import "testing"

func SimpleTest(t *testing.T) {
	amount := 500.0
	expectedTax := 5.0

	actual := CalculateTax(amount)

	if actual != expectedTax {
		t.Errorf("Expected %f but got %f", expectedTax, actual)
	}
}

func TestTable(t *testing.T) {
	tt := map[string]struct {
		amount      float64
		expectedTax float64
	}{
		"calculate tax when amount less than 1000": {
			amount:      500,
			expectedTax: 5.0,
		},
		"calculate tax when amount greater than 1000": {
			amount:      1001,
			expectedTax: 10.0,
		},
	}

	for _, tc := range tt {
		actual := CalculateTax(tc.amount)

		if actual != tc.expectedTax {
			t.Errorf("Expected %f but got %f", tc.expectedTax, actual)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
