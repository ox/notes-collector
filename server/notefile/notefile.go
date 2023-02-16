package notefile

import (
	"encoding/json"
	"fmt"
	"os"
)

type Notefile struct {
	Notes []*Note
}

func NewNotefile() *Notefile {
	return &Notefile{
		Notes: make([]*Note, 0),
	}
}

func Load(localpath string) (*Notefile, error) {
	b, err := os.ReadFile(localpath)
	if err != nil {
		return nil, fmt.Errorf("could not read notefile: %w", err)
	}

	nf := NewNotefile()
	if err = json.Unmarshal(b, nf); err != nil {
		return nil, fmt.Errorf("could not parse notefile: %w", err)
	}

	return nf, nil
}

func (n *Notefile) Save(path string) error {
	b, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return fmt.Errorf("could not save notefile: %w", err)
	}

	if err = os.WriteFile(path, b, 0755); err != nil {
		return fmt.Errorf("could not write notefile: %w", err)
	}

	return nil
}
