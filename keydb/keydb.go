package keydb

// START OMIT
type DB struct {
	data map[string]interface{}
}

func NewDB() *DB {
	d := &DB{}
	d.data = make(map[string]interface{})
	return d
}

func (d *DB) Get(key string) (interface{}, bool) {
	v, ok := d.data[key]
	return v, ok
}

func (d *DB) Set(key string, value interface{}) bool {
	_, existed := d.data[key]
	d.data[key] = value
	return existed
}

// END OMIT
