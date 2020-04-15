package cyoa

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StoryEntry struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ReadStoryJSON(in []byte) (map[string]StoryEntry, error){
	story := make(map[string]StoryEntry)

	err := json.Unmarshal(in, &story)
	if err != nil {
		return nil, err
	}

	//check that intro exists and all chapter refs are valid
	if _, introFound := story["intro"]; introFound == false {
		return nil, errors.New("intro not found")
	}

	for _, se := range story {
		for _, opt := range se.Options {
			if _, arcFound := story[opt.Arc]; arcFound == false {
				return nil, fmt.Errorf("arc %s not found in story", opt.Arc)
			}
		}
	}


	return story, nil
}

func (e *StoryEntry) String() string {
	var strs []string
	strs = append(strs, e.Title, "\n\n")
	strs = append(strs, strings.Join(e.Story, "\n"), "\n\n")
	for i, opt := range e.Options {
		strs = append(strs, fmt.Sprintf("%d: %s\n", i+1, opt.Text))
	}
	strs = append(strs, "---------\n")

	return strings.Join(strs, "")
}