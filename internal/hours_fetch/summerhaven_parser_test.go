package hoursfetch

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/hoodnoah/ifw_ranges/internal/types"
)

const dataPath = "../../testdata/summerhaven.html"

func loadExampleValues() types.RangeHours {
	return types.RangeHours{
		Range: "Summerhaven",
		Hours: []types.RangeDayHours{
			types.NewRangeHours(2025, time.May, 28, 10, 13),
			types.NewRangeHours(2025, time.May, 29, 10, 13),
			types.NewRangeHours(2025, time.May, 30, 10, 13),
			types.NewRangeHours(2025, time.May, 31, 10, 17),
			types.NewRangeHours(2025, time.June, 2, 10, 17),
			types.NewRangeHours(2025, time.June, 3, 10, 17),
			types.NewRangeHours(2025, time.June, 4, 10, 17),
			types.NewRangeHours(2025, time.June, 5, 10, 17),
			types.NewRangeHours(2025, time.June, 6, 10, 17),
			types.NewRangeHours(2025, time.June, 7, 10, 13),
			types.NewRangeHours(2025, time.June, 8, 10, 17),
			types.NewRangeHours(2025, time.June, 9, 10, 17),
			types.NewRangeHours(2025, time.June, 10, 10, 17),
			types.NewRangeHours(2025, time.June, 11, 10, 17),
			types.NewRangeHours(2025, time.June, 12, 10, 13),
			types.NewRangeHours(2025, time.June, 13, 10, 13),
			types.NewRangeHours(2025, time.June, 14, 10, 17),
			types.NewRangeHours(2025, time.June, 15, 10, 17),
			types.NewRangeHours(2025, time.June, 16, 10, 13),
			types.NewRangeHours(2025, time.June, 17, 10, 17),
			types.NewRangeHours(2025, time.June, 18, 10, 17),
			types.NewRangeHours(2025, time.June, 19, 10, 13),
			types.NewRangeHours(2025, time.June, 20, 10, 13),
			types.NewRangeHours(2025, time.June, 21, 10, 13),
			types.NewRangeHours(2025, time.June, 22, 10, 13),
			types.NewRangeHours(2025, time.June, 23, 10, 17),
			types.NewRangeHours(2025, time.June, 24, 10, 17),
			types.NewRangeHours(2025, time.June, 25, 10, 13),
			types.NewRangeHours(2025, time.June, 26, 10, 17),
			types.NewRangeHours(2025, time.June, 27, 10, 17),
			types.NewRangeHours(2025, time.June, 28, 10, 17),
			types.NewRangeHours(2025, time.June, 29, 10, 13),
			types.NewRangeHours(2025, time.June, 30, 10, 13),
		},
	}
}

func TestSummerHavenParse(t *testing.T) {
	// read test data file into bytes
	testFile, err := os.ReadFile(dataPath)

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

	t.Run("parses all examples correctly and in the correct order", func(t *testing.T) {
		expected := loadExampleValues()
		loc, _ := time.LoadLocation("America/New_York")
		fetchTime := time.Date(2025, time.June, 9, 17, 14, 0, 0, loc)

		file, err := os.ReadFile(dataPath)
		if err != nil {
			t.Fatalf("failed to load example file %v", err)
		}

		parser := NewSummerHavenParser()

		actual, err := parser.ParseDatesFromHtml(file, fetchTime)
		if err != nil {
			t.Fatalf("expected parser not to error on well-formed example file %v", err)
		}

		for i, res := range actual.Hours {
			if !reflect.DeepEqual(expected.Hours[i], res) {
				t.Fatalf("expected the following to equal, but didn't: \n\n Expected: \n %v  \n\n Actual: \n %v", expected.Hours[i], res)
			}
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
