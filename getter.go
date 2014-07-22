package cookoo

import (
	"reflect"
)

type Getter interface {
	Get(string, interface{}) ContextValue
	Has(string) (ContextValue, bool)
}

// GettableDatasource Makes a KeyValueDatasource match the Getter interface.
//
// In future versions of Cookoo, core Datasources will directly implement Getter.
type GettableDatasource struct {
	KeyValueDatasource
}

func (g *GettableDatasource) Get(key string, defaultVal interface{}) ContextValue {
	ret := g.KeyValueDatasource.Value(key)
	if ret == nil || !reflect.ValueOf(ret).IsValid() {
		return defaultVal
	}
	return ret
}

func (g *GettableDatasource) Has(key string) (ContextValue, bool) {
	ret := g.KeyValueDatasource.Value(key)
	if ret == nil || !reflect.ValueOf(ret).IsValid() {
		return nil, false
	}
	return ret, true
}

// GetString is a convenience function for getting strings.
//
// This simplifies getting strings from a Context, a Params, or a
// GettableDatasource.
func GetString(key, defaultValue string, source Getter) string {
	return source.Get(key, defaultValue).(string)
}

func GetBool(key string, defaultValue bool, source Getter) bool {
	return source.Get(key, defaultValue).(bool)
}

func GetInt(key string, defaultValue int, source Getter) int {
	return source.Get(key, defaultValue).(int)
}

func GetInt64(key string, defaultValue int64, source Getter) int64 {
	return source.Get(key, defaultValue).(int64)
}

func GetInt32(key string, defaultValue int32, source Getter) int32 {
	return source.Get(key, defaultValue).(int32)
}

func GetUint64(key string, defaultVal uint64, source Getter) uint64 {
	return source.Get(key, defaultVal).(uint64)
}

func GetFloat64(key string, defaultVal float64, source Getter) float64 {
	return source.Get(key, defaultVal).(float64)
}

// HasString is a convenience function to perform Has() and return a string.
func HasString(key string, source Getter) (string, bool) {
	v, ok := source.Has(key)
	if !ok {
		return "", ok
	}
	strval, kk := v.(string)
	if !kk {
		return "", kk
	}
	return strval, kk
}

func HasBool(key string, defaultValue bool, source Getter) (bool, bool) {
	v, ok := source.Has(key)
	if !ok {
		return false, ok
	}
	strval, kk := v.(bool)
	if !kk {
		return false, kk
	}
	return strval, kk
}

func HasInt(key string, defaultValue int, source Getter) (int, bool) {
	v, ok := source.Has(key)
	if !ok {
		return 0, ok
	}
	val, kk := v.(int)
	if !kk {
		return 0, kk
	}
	return val, kk
}

func HasInt64(key string, defaultValue int64, source Getter) (int64, bool) {
	v, ok := source.Has(key)
	if !ok {
		return 0, ok
	}
	val, kk := v.(int64)
	if !kk {
		return 0, kk
	}
	return val, kk
}

func HasInt32(key string, defaultValue int32, source Getter) (int32, bool) {
	v, ok := source.Has(key)
	if !ok {
		return 0, ok
	}
	val, kk := v.(int32)
	if !kk {
		return 0, kk
	}
	return val, kk
}

func HasUint64(key string, defaultVal uint64, source Getter) (uint64, bool) {
	v, ok := source.Has(key)
	if !ok {
		return 0, ok
	}
	val, kk := v.(uint64)
	if !kk {
		return 0, kk
	}
	return val, kk
}

func HasFloat64(key string, defaultVal float64, source Getter) (float64, bool) {
	v, ok := source.Has(key)
	if !ok {
		return 0, ok
	}
	val, kk := v.(float64)
	if !kk {
		return 0, kk
	}
	return val, kk
}

// GetFromFirst gets the value from the first Getter that has the key.
//
// If no Getter has the key, the default value is returned, and the returned
// Getter is an instance of DefaultGetter.
func GetFromFirst(key string, defaultVal interface{}, sources ...Getter) (ContextValue, Getter) {
	for _, s := range sources {
		val, ok := s.Has(key)
		if ok {
			return val, s
		}
	}

	return defaultVal, &DefaultGetter{defaultVal}
}

// DefaultGetter represents a Getter instance for a default value.
//
// A default getter always returns the given default value.
type DefaultGetter struct {
	val ContextValue
}

func (e *DefaultGetter) Get(name string, value interface{}) ContextValue {
	return e.val
}
func (e *DefaultGetter) Has(name string) (ContextValue, bool) {
	return e.val, true
}