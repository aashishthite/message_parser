package parse

type Parser interface {
	Parse(string) (*ParseResult, error)
}

func NewParser() Parser {
	return &parseImpl{}
}

type parseImpl struct {
}

func (p *parseImpl) Parse(msg string) (*ParseResult, error) {

	parsedMentions, err := parseMentions(msg)
	if err != nil {
		return nil, err
	}
	if len(parsedMentions) == 0 {
		parsedMentions = nil
	}

	parsedEmoticons, err := parseEmoticons(msg)
	if err != nil {
		return nil, err
	}
	if len(parsedEmoticons) == 0 {
		parsedEmoticons = nil
	}

	parsedLinks, err := parseLinks(msg)
	if err != nil {
		return nil, err
	}
	if len(parsedLinks) == 0 {
		parsedLinks = nil
	}

	retval := &ParseResult{
		Mentions:  parsedMentions,
		Emoticons: parsedEmoticons,
		Links:     parsedLinks,
	}
	return retval, nil
}
