package parse

import "regexp"

const mentionsRegexp = `@\w+`

func parseMentions(msg string) ([]string, error) {
	re := regexp.MustCompile(mentionsRegexp)
	retval := make([]string, 0)

	for {
		match := re.FindStringIndex(msg)
		if match != nil {
			retval = append(retval, msg[match[0]+1:match[1]])
			msg = msg[match[1]:len(msg)]
		} else {
			return retval, nil
		}

	}
}
