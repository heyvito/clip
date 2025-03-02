// Code generated by generators/cli; DO NOT EDIT.

package clip

import "fmt"

func (c *CLI) Int(name string) int {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return int(fv)
	case int8:
		return int(fv)
	case int16:
		return int(fv)
	case int32:
		return int(fv)
	case int64:
		return int(fv)
	case uint:
		return int(fv)
	case uint8:
		return int(fv)
	case uint16:
		return int(fv)
	case uint32:
		return int(fv)
	case uint64:
		return int(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as int: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchInt(name string) (int, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return int(fv), true
	case int8:
		return int(fv), true
	case int16:
		return int(fv), true
	case int32:
		return int(fv), true
	case int64:
		return int(fv), true
	case uint:
		return int(fv), true
	case uint8:
		return int(fv), true
	case uint16:
		return int(fv), true
	case uint32:
		return int(fv), true
	case uint64:
		return int(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as int: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Int8(name string) int8 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return int8(fv)
	case int8:
		return int8(fv)
	case int16:
		return int8(fv)
	case int32:
		return int8(fv)
	case int64:
		return int8(fv)
	case uint:
		return int8(fv)
	case uint8:
		return int8(fv)
	case uint16:
		return int8(fv)
	case uint32:
		return int8(fv)
	case uint64:
		return int8(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as int8: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchInt8(name string) (int8, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return int8(fv), true
	case int8:
		return int8(fv), true
	case int16:
		return int8(fv), true
	case int32:
		return int8(fv), true
	case int64:
		return int8(fv), true
	case uint:
		return int8(fv), true
	case uint8:
		return int8(fv), true
	case uint16:
		return int8(fv), true
	case uint32:
		return int8(fv), true
	case uint64:
		return int8(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as int8: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Int16(name string) int16 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return int16(fv)
	case int8:
		return int16(fv)
	case int16:
		return int16(fv)
	case int32:
		return int16(fv)
	case int64:
		return int16(fv)
	case uint:
		return int16(fv)
	case uint8:
		return int16(fv)
	case uint16:
		return int16(fv)
	case uint32:
		return int16(fv)
	case uint64:
		return int16(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as int16: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchInt16(name string) (int16, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return int16(fv), true
	case int8:
		return int16(fv), true
	case int16:
		return int16(fv), true
	case int32:
		return int16(fv), true
	case int64:
		return int16(fv), true
	case uint:
		return int16(fv), true
	case uint8:
		return int16(fv), true
	case uint16:
		return int16(fv), true
	case uint32:
		return int16(fv), true
	case uint64:
		return int16(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as int16: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Int32(name string) int32 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return int32(fv)
	case int8:
		return int32(fv)
	case int16:
		return int32(fv)
	case int32:
		return int32(fv)
	case int64:
		return int32(fv)
	case uint:
		return int32(fv)
	case uint8:
		return int32(fv)
	case uint16:
		return int32(fv)
	case uint32:
		return int32(fv)
	case uint64:
		return int32(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as int32: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchInt32(name string) (int32, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return int32(fv), true
	case int8:
		return int32(fv), true
	case int16:
		return int32(fv), true
	case int32:
		return int32(fv), true
	case int64:
		return int32(fv), true
	case uint:
		return int32(fv), true
	case uint8:
		return int32(fv), true
	case uint16:
		return int32(fv), true
	case uint32:
		return int32(fv), true
	case uint64:
		return int32(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as int32: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Int64(name string) int64 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return int64(fv)
	case int8:
		return int64(fv)
	case int16:
		return int64(fv)
	case int32:
		return int64(fv)
	case int64:
		return int64(fv)
	case uint:
		return int64(fv)
	case uint8:
		return int64(fv)
	case uint16:
		return int64(fv)
	case uint32:
		return int64(fv)
	case uint64:
		return int64(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as int64: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchInt64(name string) (int64, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return int64(fv), true
	case int8:
		return int64(fv), true
	case int16:
		return int64(fv), true
	case int32:
		return int64(fv), true
	case int64:
		return int64(fv), true
	case uint:
		return int64(fv), true
	case uint8:
		return int64(fv), true
	case uint16:
		return int64(fv), true
	case uint32:
		return int64(fv), true
	case uint64:
		return int64(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as int64: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Uint(name string) uint {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return uint(fv)
	case int8:
		return uint(fv)
	case int16:
		return uint(fv)
	case int32:
		return uint(fv)
	case int64:
		return uint(fv)
	case uint:
		return uint(fv)
	case uint8:
		return uint(fv)
	case uint16:
		return uint(fv)
	case uint32:
		return uint(fv)
	case uint64:
		return uint(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as uint: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchUint(name string) (uint, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return uint(fv), true
	case int8:
		return uint(fv), true
	case int16:
		return uint(fv), true
	case int32:
		return uint(fv), true
	case int64:
		return uint(fv), true
	case uint:
		return uint(fv), true
	case uint8:
		return uint(fv), true
	case uint16:
		return uint(fv), true
	case uint32:
		return uint(fv), true
	case uint64:
		return uint(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as uint: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Uint8(name string) uint8 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return uint8(fv)
	case int8:
		return uint8(fv)
	case int16:
		return uint8(fv)
	case int32:
		return uint8(fv)
	case int64:
		return uint8(fv)
	case uint:
		return uint8(fv)
	case uint8:
		return uint8(fv)
	case uint16:
		return uint8(fv)
	case uint32:
		return uint8(fv)
	case uint64:
		return uint8(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as uint8: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchUint8(name string) (uint8, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return uint8(fv), true
	case int8:
		return uint8(fv), true
	case int16:
		return uint8(fv), true
	case int32:
		return uint8(fv), true
	case int64:
		return uint8(fv), true
	case uint:
		return uint8(fv), true
	case uint8:
		return uint8(fv), true
	case uint16:
		return uint8(fv), true
	case uint32:
		return uint8(fv), true
	case uint64:
		return uint8(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as uint8: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Uint16(name string) uint16 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return uint16(fv)
	case int8:
		return uint16(fv)
	case int16:
		return uint16(fv)
	case int32:
		return uint16(fv)
	case int64:
		return uint16(fv)
	case uint:
		return uint16(fv)
	case uint8:
		return uint16(fv)
	case uint16:
		return uint16(fv)
	case uint32:
		return uint16(fv)
	case uint64:
		return uint16(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as uint16: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchUint16(name string) (uint16, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return uint16(fv), true
	case int8:
		return uint16(fv), true
	case int16:
		return uint16(fv), true
	case int32:
		return uint16(fv), true
	case int64:
		return uint16(fv), true
	case uint:
		return uint16(fv), true
	case uint8:
		return uint16(fv), true
	case uint16:
		return uint16(fv), true
	case uint32:
		return uint16(fv), true
	case uint64:
		return uint16(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as uint16: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Uint32(name string) uint32 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return uint32(fv)
	case int8:
		return uint32(fv)
	case int16:
		return uint32(fv)
	case int32:
		return uint32(fv)
	case int64:
		return uint32(fv)
	case uint:
		return uint32(fv)
	case uint8:
		return uint32(fv)
	case uint16:
		return uint32(fv)
	case uint32:
		return uint32(fv)
	case uint64:
		return uint32(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as uint32: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchUint32(name string) (uint32, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return uint32(fv), true
	case int8:
		return uint32(fv), true
	case int16:
		return uint32(fv), true
	case int32:
		return uint32(fv), true
	case int64:
		return uint32(fv), true
	case uint:
		return uint32(fv), true
	case uint8:
		return uint32(fv), true
	case uint16:
		return uint32(fv), true
	case uint32:
		return uint32(fv), true
	case uint64:
		return uint32(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as uint32: It's a %T", f.name, f.value))
	}
}

func (c *CLI) Uint64(name string) uint64 {
	f := c.find(name)
	if f == nil {
		return 0
	}
	switch fv := f.value.(type) {
	case nil:
		return 0
	case int:
		return uint64(fv)
	case int8:
		return uint64(fv)
	case int16:
		return uint64(fv)
	case int32:
		return uint64(fv)
	case int64:
		return uint64(fv)
	case uint:
		return uint64(fv)
	case uint8:
		return uint64(fv)
	case uint16:
		return uint64(fv)
	case uint32:
		return uint64(fv)
	case uint64:
		return uint64(fv)
	default:
		panic(fmt.Sprintf("Cannot read %s as uint64: It's a %T", f.name, f.value))
	}
}

func (c *CLI) FetchUint64(name string) (uint64, bool) {
	f := c.find(name)
	if f == nil {
		return 0, false
	}
	switch fv := f.value.(type) {
	case nil:
		return 0, false
	case int:
		return uint64(fv), true
	case int8:
		return uint64(fv), true
	case int16:
		return uint64(fv), true
	case int32:
		return uint64(fv), true
	case int64:
		return uint64(fv), true
	case uint:
		return uint64(fv), true
	case uint8:
		return uint64(fv), true
	case uint16:
		return uint64(fv), true
	case uint32:
		return uint64(fv), true
	case uint64:
		return uint64(fv), true
	default:
		panic(fmt.Sprintf("Cannot read %s as uint64: It's a %T", f.name, f.value))
	}
}
