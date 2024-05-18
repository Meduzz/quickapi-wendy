package api

import "encoding/json"

type (
	Create struct {
		Entity json.RawMessage
	}

	Read struct {
		ID      string
		Filters map[string]map[string]string
	}

	Update struct {
		ID      string
		Entity  json.RawMessage
		Filters map[string]map[string]string
	}

	Delete struct {
		ID      string
		Filters map[string]map[string]string
	}

	Search struct {
		Skip    int
		Take    int
		Where   map[string]string
		Filters map[string]map[string]string
	}

	Patch struct {
		ID      string
		Data    map[string]any
		Filters map[string]map[string]string
	}
)
