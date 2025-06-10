package hoursfetch

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hoodnoah/ifw_ranges/internal/types"
)

type CleanedLine string

type SummerHavenParser struct {
	scheduleLinesRegex regexp.Regexp
	dayOfWeekRegex     regexp.Regexp
}

var timeExtractRegex = regexp.MustCompile(`^(\d{1,2})(?::(\d{2}))?m$`)

func NewSummerHavenParser() SummerHavenParser {
	return SummerHavenParser{
		scheduleLinesRegex: *regexp.MustCompile(`(?m)^.*\d+[ap]m\s*.*<br>`),
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
	line = strings.ReplaceAll(line, "<ul>", "")
	line = strings.ReplaceAll(line, "</ul>", "")
	line = strings.ReplaceAll(line, "<br>", "")
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "-")
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

func tryParseTime(timeStr string) (time.Time, error) {
	cleaned := normalizeTime(timeStr)
	layouts := []string{"3pm", "3:04pm"}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, cleaned); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse time %q", timeStr)
}

func normalizeTime(input string) string {
	s := strings.ToLower(strings.TrimSpace(input))

	// Fix common case: trailing "m" without "a/p"
	if matches := timeExtractRegex.FindStringSubmatch(s); matches != nil {
		hourStr := matches[1]
		hour, _ := strconv.Atoi(hourStr)

		// heuristic; assume AM only for 9, 10 or 11, PM otherwise
		if hour >= 9 && hour <= 11 {
			if matches[2] != "" {
				return hourStr + ":" + matches[2] + "am"
			}
			return hourStr + "am"
		} else {
			if matches[2] != "" {
				return hourStr + ":" + matches[2] + "pm"
			}
			return hourStr + "pm"
		}
	}

	return s
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

	start, err := tryParseTime(startPart)
	if err != nil {
		return nil, fmt.Errorf("invalid start time %q: %w", startPart, err)
	}

	end, err := tryParseTime(endPart)
	if err != nil {
		return nil, fmt.Errorf("invalid end time %q: %w", endPart, err)
	}

	// handle mistakenly swapped AM/PM suffixes
	if start.After(end) {
		start = start.Add(-12 * time.Hour)
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
