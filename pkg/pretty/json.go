package pretty

import (
	"encoding/json"
	"io"
)

// JSON write pretty json
func JSON(w io.Writer, jsonData interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(jsonData)
}
