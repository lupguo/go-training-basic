package sub

import "errors"

type Slice []interface{}

var elem_not_exist = errors.New("elem not exist")

func (s *Slice) Remove(val interface{}) (slice *Slice, err error) {
	for i, v := range *s {
		if isEqual(v, val) {
			if i == len(*s)-1 {
				*s = (*s)[:i]
			} else {
				*s = append((*s)[:i], (*s)[i+1:]...) // be careful (*s)[i+1:]...|(*s)[i+1]|(*s)[i+2:]
			}
			return s, nil
		}
	}
	return s, elem_not_exist
}

func isEqual(a, b interface{}) bool {
	return a == b
}
