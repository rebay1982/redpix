package redpix

import (
	"testing"
)

type windowConfigValidateTest struct {
	width, height int
	expected bool
}

var windowConfigValidateTests = []windowConfigValidateTest {
	windowConfigValidateTest{0, 0, false},
	windowConfigValidateTest{100, 0, false},
	windowConfigValidateTest{1981, 100, false},
	windowConfigValidateTest{100, 1081, false},
	windowConfigValidateTest{1981, 1081, false},
	windowConfigValidateTest{100, 100, true},
}

func TestWindowConfigValidate(t *testing.T) {
	for ti, test := range windowConfigValidateTests {
		config := WindowConfig{Width: test.width, Height: test.height}

		if result := config.validate(); result != test.expected {
			t.Errorf("Scenario %d: Expected [%t], got [%t] for width=[%d] and height=[%d]",
				ti, test.expected, result, test.width, test.height)
		}
	}
}
