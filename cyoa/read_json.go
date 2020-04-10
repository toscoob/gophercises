package cyoa

import(
	"encoding/json"
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

	//TODO check for intro chapter
	//TODO check that all chapter refs are valid

	return story, nil
}
