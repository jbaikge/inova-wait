package main

import (
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetHTML(url string) (html []byte, err error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// All the data is on one long line, spaced out for reference below:
//
// <div class='er-location'>
//   <a href='[redacted]' target='_blank'>
//     <h3>43<abbr title='minute' class='min-abb'>min</abbr></h3>
//     <div class='location-details'>
//       <span class='location-name'>Inova Alexandria Hospital</span>
//       <div class='location-address'>
//         <span class='location-icon'></span>
//         <span class='location-address-line1'>4320 Seminary Rd</span>
//         <span class='location-address-line2'>Alexandria, VA 22304</span>
//       </div>
//     </div>
//     <div class='clearer'></div>
//   </a>
// </div>
func ParseWaits(html []byte) (waits []Wait, err error) {
	re := regexp.MustCompile(`(?m)<div class='er-location'>.*?<h3>(\d+).*<span class='location-name'>([^<]+)</span>.*</div>`)
	matches := re.FindAllSubmatch(html, 16)
	waits = make([]Wait, len(matches))
	for i, match := range matches {
		waits[i].Id = i + 1
		if waits[i].Minutes, err = strconv.Atoi(string(match[1])); err != nil {
			return
		}
		loc := strings.Replace(string(match[2]), "&mdash;", "-", 1)
		waits[i].Location = loc
	}
	return
}
