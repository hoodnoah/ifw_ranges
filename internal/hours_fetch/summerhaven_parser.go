package hoursfetch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/hoodnoah/ifw_ranges/internal/types"
)

type CleanedLine string

type SummerHavenParser struct {
	scheduleLinesRegex regexp.Regexp
	dayOfWeekRegex     regexp.Regexp
}

func NewSummerHavenParser() SummerHavenParser {
	return SummerHavenParser{
		scheduleLinesRegex: *regexp.MustCompile(`(?m)^.*\d+[ap]m\s*<br>`),
		dayOfWeekRegex:     *regexp.MustCompile(`(?m)^.*,`),
	}
}

// Parses, from an html file (read into bytes), the hours for the Summerhaven Range.
func (s *SummerHavenParser) ParseDatesFromHtml(webpage []byte, fetchedAt time.Time) (*types.RangeHours, error) {
	lines, err := s.ExtractScheduleLines(webpage)
	if err != nil {
		return nil, err
	}

	cleanedLines, err := cleanLines(lines)
	if err != nil {
		return nil, err
	}

	parsedLines, err := parseCleanedLines(cleanedLines)
	if err != nil {
		return nil, err
	}

	resolvedDates := resolveYears(parsedLines, fetchedAt)

	return &types.RangeHours{
		Range: "Summerhaven",
		Hours: resolvedDates,
	}, nil
}

func (s *SummerHavenParser) ExtractScheduleLines(pageBytes []byte) ([]string, error) {
	html := string(pageBytes)

	lines := s.scheduleLinesRegex.FindAllString(html, -1) // -1 -> all matches
	if len(lines) < 1 {
		return nil, errors.New("no schedule lines found")
	}
	return lines, nil
}

// cleans schedule lines, e.g.:
// <p>Wednesday, May 28 - 10am - 1pm <br> -> Wednesday, May 28 - 10am - 1pm
func cleanLine(line string) (CleanedLine, error) {
	var normalizeHyphens = regexp.MustCompile(`\s*-\s*`)
	line = strings.ReplaceAll(line, "<p>", "")
	line = strings.ReplaceAll(line, "<br>", "")
	line = strings.TrimSpace(line)
	line = normalizeHyphens.ReplaceAllString(line, " - ")

	return CleanedLine(line), nil
}

func cleanLines(lines []string) ([]CleanedLine, error) {
	var cleanLines []CleanedLine
	for _, line := range lines {
		cleaned, err := cleanLine(line)
		if err != nil {
			return nil, err
		}
		cleanLines = append(cleanLines, cleaned)
	}

	return cleanLines, nil
}

func parseCleanedLine(cleanedLine CleanedLine) (*types.PartialRangeDayHours, error) {
	// strip off day-of-week portion
	parts := strings.SplitN(string(cleanedLine), ",", 2)
	if len(parts) != 2 {
		return nil, errors.New("expected a single comma in schedule line")
	}
	rest := strings.TrimSpace(parts[1])

	parts = strings.SplitN(rest, " - ", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("expected 3 parts in schedule line, got: %q", len(parts))
	}

	datePart := parts[0]
	startPart := parts[1]
	endPart := parts[2]

	date, err := time.Parse("January 2", datePart)
	if err != nil {
		return nil, fmt.Errorf("invalid date %q: %w", datePart, err)
	}

	start, err := time.Parse("3pm", startPart)
	if err != nil {
		return nil, fmt.Errorf("invalid start time %q: %w", startPart, err)
	}

	end, err := time.Parse("3pm", endPart)
	if err != nil {
		return nil, fmt.Errorf("invalid end time %q: %w", endPart, err)
	}

	return &types.PartialRangeDayHours{
		Month:     date.Month(),
		Day:       date.Day(),
		StartHour: start.Hour(),
		StartMin:  start.Minute(),
		EndHour:   end.Hour(),
		EndMin:    end.Minute(),
	}, nil
}

func parseCleanedLines(cleanedLines []CleanedLine) ([]*types.PartialRangeDayHours, error) {

	var parsedCleanedLines []*types.PartialRangeDayHours
	for _, line := range cleanedLines {
		parsedLine, err := parseCleanedLine(line)
		if err != nil {
			return nil, err
		}
		parsedCleanedLines = append(parsedCleanedLines, parsedLine)
	}

	return parsedCleanedLines, nil
}

func resolveYears(partials []*types.PartialRangeDayHours, ref time.Time) []types.RangeDayHours {
	var resolved []types.RangeDayHours
	for _, partial := range partials {
		resolved = append(resolved, partial.ResolveYear(ref))
	}

	return resolved
}
