package hoursfetch

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hoodnoah/ifw_ranges/internal/types"
	"github.com/hoodnoah/ifw_ranges/testdata/summerhaven"
)

const dataPath = "../../testdata/summerhaven/html/summerhaven_06_09_2025.html"

var loc, _ = time.LoadLocation("America/New_York")

func expectedFor(filename string) types.RangeHours {
	switch filename {
	case "summerhaven_08_03_2023.html":
		return types.RangeHours{
			Range: "Summerhaven",
			Hours: []types.RangeDayHours{
				{
					Date:      time.Date(2023, time.July, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.July, 29, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.July, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.July, 31, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 1, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 2, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 3, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 13, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 4, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 6, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 7, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 8, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 9, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 10, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 13, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 11, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 13, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 14, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 15, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 16, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 17, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 18, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 20, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 21, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 22, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 23, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 24, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 25, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 26, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 29, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.August, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				},
			},
		}
	case "summerhaven_11_28_2023.html":
		return types.RangeHours{
			Range: "Summerhaven",
			Hours: []types.RangeDayHours{
				{
					Date:      time.Date(2023, time.November, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 13, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.November, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.November, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 1, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 2, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 3, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 4, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 5, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 7, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 8, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 10, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 11, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 15, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 16, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 17, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 18, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 19, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 21, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 16, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 22, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 29, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				}, {
					Date:      time.Date(2023, time.December, 31, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
				},
			},
		}
	case "summerhaven_06_23_2024.html":
		return types.RangeHours{
			Range: "Summerhaven",
			Hours: []types.RangeDayHours{
				{
					Date:      time.Date(2024, time.May, 31, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 2, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 3, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 4, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 5, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 7, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 8, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 9, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 10, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 11, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 12, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 13, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 14, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 15, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 16, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 17, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 22, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 23, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 24, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 25, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 26, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				},
			},
		}
	case "summerhaven_06_09_2025.html":
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
	case "summerhaven_07_16_2024.html":
		return types.RangeHours{
			Range: "Summerhaven",
			Hours: []types.RangeDayHours{
				{
					Date:      time.Date(2024, time.June, 26, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.June, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.June, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 1, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 2, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 5, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 7, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 8, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 9, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 10, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 11, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 12, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 14, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 15, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 16, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 17, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 18, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 19, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 20, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 21, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 22, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 23, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 24, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 13, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 25, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 13, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 26, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 27, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 17, Minute: 0},
				}, {
					Date:      time.Date(2024, time.July, 28, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 29, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 30, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				}, {
					Date:      time.Date(2024, time.July, 31, 0, 0, 0, 0, loc),
					StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
					EndTime:   types.TimeOfDay{Hour: 13, Minute: 30},
				},
			},
		}
	default:
		panic("no expected data for " + filename)
	}
}

func TestSummerHavenParse(t *testing.T) {
	// read test data file into bytes
	// testFile, err := os.ReadFile(dataPath)

	// if err != nil {
	// 	t.Fatalf("failed to read the summerhaven test data %v", err)
	// }

	loc, _ := time.LoadLocation("America/New_York")

	// dateFetched := time.Date(2025, 6, 9, 17, 14, 0, 0, loc)

	// t.Run("parses the first date correctly", func(t *testing.T) {
	// 	loc, _ := time.LoadLocation("America/New_York")

	// 	expectedFirstDate := types.RangeDayHours{
	// 		Date:      time.Date(2025, time.May, 28, 0, 0, 0, 0, loc),
	// 		StartTime: types.TimeOfDay{Hour: 10, Minute: 0},
	// 		EndTime:   types.TimeOfDay{Hour: 13, Minute: 0},
	// 	}

	// 	parser := NewSummerHavenParser()

	// 	result, err := parser.ParseDatesFromHtml(testFile, dateFetched)
	// 	if err != nil {
	// 		t.Fatalf("expected to receive no error, received one %v", err)
	// 	}
	// 	if result == nil {
	// 		t.Fatalf("expected a non-null result, received null")
	// 	}

	// 	if !reflect.DeepEqual(expectedFirstDate, result.Hours[0]) {
	// 		t.Fatalf("expected to receive %v, received %v", expectedFirstDate, result.Hours[0])
	// 	}
	// })

	for _, filename := range summerhaven.ListFileNames() {
		t.Run(filename, func(t *testing.T) {
			html, err := summerhaven.LoadHTML(filename)
			if err != nil {
				t.Fatalf("could not read testfile %s: %v", filename, err)
			}

			// Extract fetch date from the filename - e.g. "summerhaven_06_09_2025.html"
			datePart := strings.TrimSuffix(strings.TrimPrefix(filename, "summerhaven_"), ".html")
			parsedTime, err := time.ParseInLocation("01_02_2006", datePart, loc)
			if err != nil {
				t.Fatalf("failed to parse date from filename: %v", err)
			}

			parser := NewSummerHavenParser()
			actual, err := parser.ParseDatesFromHtml(html, parsedTime)
			if err != nil {
				t.Fatalf("parser error on %s: %v", filename, err)
			}

			expected := expectedFor(filename)

			if expected.Range != actual.Range {
				t.Fatalf("expected Range value to be %s, received %s", expected.Range, actual.Range)
			}

			for i, act := range actual.Hours {
				if !reflect.DeepEqual(expected.Hours[i], act) {
					t.Errorf("mismatch in parsed data for %s:\n\nExpected:\n%+v\n\nActual:\n%+v", filename, expected.Hours[i], act)
				}
			}
		})
	}

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
