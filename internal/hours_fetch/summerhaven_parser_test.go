package hoursfetch

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/hoodnoah/ifw_ranges/internal/types"
)

type MockTimeHelper struct{}

func (m *MockTimeHelper) GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("America/New_York")
	testTime := time.Date(2025, time.June, 9, 17, 22, 0, 0, loc)

	return testTime
}

func TestSummerHavenParse(t *testing.T) {
	// read test data file into bytes
	testFile, err := os.ReadFile("../../testdata/summerhaven.html")

	if err != nil {
		t.Fatalf("failed to read the summerhaven test data %v", err)
	}

	loc, _ := time.LoadLocation("America/New_York")

	dateFetched := time.Date(2025, 6, 9, 17, 14, 0, 0, loc)

	t.Run("parses the first date correctly", func(t *testing.T) {
		loc, _ := time.LoadLocation("America/New_York")

		expectedFirstDate := types.RangeDayHours{
			Date:      time.Date(2025, time.May, 28, 0, 0, 0, 0, loc),
			StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
			EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
		}

		parser := NewSummerHavenParser()

		result, err := parser.ParseDatesFromHtml(testFile, dateFetched)
		if err != nil {
			t.Fatalf("expected to receive no error, received one %v", err)
		}
		if result == nil {
			t.Fatalf("expected a non-null result, received null")
		}

		if !reflect.DeepEqual(expectedFirstDate, result.Hours[0]) {
			t.Fatalf("expected to receive %v, received %v", expectedFirstDate, result.Hours[0])
		}

	})
}

func TestCleanLine(t *testing.T) {
	t.Run("cleans the first line as expected", func(t *testing.T) {
		testLine := "<p>Wednesday, May 28 - 10am - 1pm <br>\n"
		expectedCleanLine := CleanedLine("Wednesday, May 28 - 10am - 1pm")

		result, err := cleanLine(testLine)
		if err != nil {
			t.Fatalf("failed to clean the first example line %v", err)
		}

		if result != expectedCleanLine {
			t.Fatalf("expected %s, received %s", expectedCleanLine, result)
		}
	})
}

func TestParseLine(t *testing.T) {
	t.Run("parses the first line as expected", func(t *testing.T) {
		testCleanedLine := CleanedLine("Wednesday, May 28 - 10am - 1pm")
		expectedResult := types.PartialRangeDayHours{
			Month:     time.May,
			Day:       28,
			StartHour: 10,
			StartMin:  0,
			EndHour:   13,
			EndMin:    0,
		}

		result, err := parseCleanedLine(testCleanedLine)

		if err != nil {
			t.Fatalf("failed to parse the example time %v", err)
		}

		if result == nil {
			t.Fatalf("expected non-null parsed date return, received nil")
		}

		if expectedResult != *result {
			t.Fatalf("expected %v, received %v", expectedResult, result)
		}
	})
}
