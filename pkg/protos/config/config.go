package config

import (
	"fmt"
	"sync"

	"github.com/Asutorufa/yuhaiin/pkg/utils/jsondb"
	"google.golang.org/protobuf/proto"
)

type DB interface {
	Batch(f ...func(*Setting) error) error
	View(f ...func(*Setting) error) error
	Dir() string
}

var _ DB = (*JsonDB)(nil)

type JsonDB struct {
	mu sync.RWMutex
	db *jsondb.DB[*Setting]
}

func NewJsonDB(path string) *JsonDB {
	s := &JsonDB{db: jsondb.Open(path, DefaultSetting(path))}
	return s
}

func (c *JsonDB) View(f ...func(*Setting) error) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, v := range f {
		if err := v(c.db.Data); err != nil {
			return err
		}
	}

	return nil
}

func (c *JsonDB) Batch(f ...func(*Setting) error) error {
	if len(f) == 0 {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cf := proto.Clone(c.db.Data).(*Setting)
	for i := range f {
		if err := f[i](cf); err != nil {
			return err
		}
	}

	if proto.Equal(c.db.Data, cf) {
		return nil
	}

	c.db.Data = cf

	if err := c.db.Save(); err != nil {
		return fmt.Errorf("save settings failed: %w", err)
	}

	return nil
}

func (c *JsonDB) Dir() string { return c.db.Dir() }
