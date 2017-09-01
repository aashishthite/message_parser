package parse

import "regexp"

const emoticonsRegexp = `\([0-9a-zA-Z]{1,15}\)`

func parseEmoticons(msg string) ([]string, error) {
	re := regexp.MustCompile(emoticonsRegexp)
	retval := make([]string, 0)

	for {
		match := re.FindStringIndex(msg)
		if match != nil {
			retval = append(retval, msg[match[0]+1:match[1]-1])
			msg = msg[match[1]:len(msg)]
		} else {
			return retval, nil
		}

	}
}
