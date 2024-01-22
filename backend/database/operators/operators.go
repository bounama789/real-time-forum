package operators

import (
	"fmt"

	"github.com/gofrs/uuid/v5"
)

func Equals(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf("=\"%v\" ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf("=\"%v\" ", v)
	} else {
		res = fmt.Sprintf("=%v ", value)
	}

	return res
}

func NotEqual(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf("!=\"%v\" ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf("!=\"%v\" ", v)
	} else {
		res = fmt.Sprintf("!=%v ", value)
	}
	return res
}

func GreaterThan(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf(">%v ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf(">%v ", v)
	} else {
		res = fmt.Sprintf(">%v ", value)
	}
	return res
}

func LessThan(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf("<%v ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf("<%v ", v)
	} else {
		res = fmt.Sprintf("<%v ", value)
	}
	return res
}

func GreaterThanOrEqual(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf(">=%v ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf(">=%v ", v)
	} else {
		res = fmt.Sprintf(">=%v ", value)
	}
	return res
}

func LessThanOrEqual(value any) string {
	var res string
	if v, ok := value.(string); ok {
		res = fmt.Sprintf("<=%v ", v)
	} else if v, ok := value.(uuid.UUID); ok {
		res = fmt.Sprintf("<=%v ", v)
	} else {
		res = fmt.Sprintf("<=%v ", value)
	}
	return res
}

func Or(key string, value string) string {
	return fmt.Sprintf(" OR %v%v ", key,value)

}

// func Or(value any) string {
// 	var res string
// 	if v, ok := value.(string); ok {
// 		res = fmt.Sprintf("OR %v ", v)
// 	} else if v, ok := value.(uuid.UUID); ok {
// 		res = fmt.Sprintf("OR %v ", v)
// 	} else {
// 		res = fmt.Sprintf("OR %v ", value)
// 	}
// 	return res
// }
