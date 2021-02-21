package common

import (
	"errors"
	"strconv"
)

func GetKey(dbName string, tableName string, fieldName string) string {
	return dbName + "_" + tableName + "_" + fieldName
}

//func GetValueString(v interface{}) string {
//	rawValue := reflect.Indirect(reflect.ValueOf(v))
//
//	stringV, _ := value2String(&rawValue)
//	return stringV
//}

// 如果早知这样，直接用xorm就好？
// 另一种类型断言
// https://www.jianshu.com/p/787cf3a41ccb
//func value2String(rawValue *reflect.Value) (str string, err error) {
//	aa := reflect.TypeOf((*rawValue).Interface())
//	vv := reflect.ValueOf((*rawValue).Interface())
//	switch aa.Kind() {
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		str = strconv.FormatInt(vv.Int(), 10)
//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//		str = strconv.FormatUint(vv.Uint(), 10)
//	case reflect.Float32, reflect.Float64:
//		str = strconv.FormatFloat(vv.Float(), 'f', -1, 64)
//	case reflect.String:
//		str = vv.String()
//	case reflect.Array, reflect.Slice:
//		switch aa.Elem().Kind() {
//		case reflect.Uint8:
//			data := rawValue.Interface().([]byte)
//			str = string(data)
//			if str == "\x00" {
//				str = "0"
//			}
//		default:
//			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
//		}
//	// time type
//	case reflect.Struct:
//		if aa.ConvertibleTo(core.TimeType) {
//			str = vv.Convert(core.TimeType).Interface().(time.Time).Format(time.RFC3339Nano)
//		} else {
//			err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
//		}
//	case reflect.Bool:
//		str = strconv.FormatBool(vv.Bool())
//	case reflect.Complex128, reflect.Complex64:
//		str = fmt.Sprintf("%v", vv.Complex())
//	/* TODO: unsupported types below
//	   case reflect.Map:
//	   case reflect.Ptr:
//	   case reflect.Uintptr:
//	   case reflect.UnsafePointer:
//	   case reflect.Chan, reflect.Func, reflect.Interface:
//	*/
//	default:
//		err = fmt.Errorf("Unsupported struct type %v", vv.Type().Name())
//	}
//	return
//}

func GetIntValue(a interface{}) (int, error) {
	switch a.(type) {
	case int8:
		return int(a.(int8)), nil
	case uint8: //byte
		return int(a.(uint8)), nil
	case int16:
		return int(a.(int16)), nil
	case uint16:
		return int(a.(uint16)), nil
	case int32: // rune
		return int(a.(int32)), nil
	case uint32:
		return int(a.(uint32)), nil
	case int64:
		return int(a.(int64)), nil
	case uint64:
		return int(a.(uint64)), nil
	case int:
		return a.(int), nil
	case uint:
		return int(a.(uint)), nil
	case float32:
		return 0, errors.New("type not support")
	case float64:
		return 0, errors.New("type not support")
	case complex64:
		return 0, errors.New("type not support")
	case complex128:
		return 0, errors.New("type not support")
	case uintptr:
		return 0, errors.New("type not support")
	case string:
		return 0, errors.New("type not support")
	default: // 其他类型有pointer， struct， array, slice ,map, interface, function, channel
		return 0, errors.New("type not support")
	}
}

func GetStringValue(a interface{}) (string, error) {
	switch a.(type) {
	case int8:
		return strconv.FormatInt(int64(a.(int8)), 10), nil
	case uint8: //byte
		return strconv.FormatUint(uint64(a.(uint8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(a.(int16)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(a.(uint16)), 10), nil
	case int32: // rune
		return strconv.FormatInt(int64(a.(int32)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(a.(uint32)), 10), nil
	case int64:
		return strconv.FormatInt(a.(int64), 10), nil
	case uint64:
		return strconv.FormatUint(a.(uint64), 10), nil
	case int:
		return strconv.Itoa(a.(int)), nil
	case uint:
		return strconv.FormatUint(uint64(a.(uint)), 10), nil
	case float32:
		return "", errors.New("type not support")
	case float64:
		return "", errors.New("type not support")
	case complex64:
		return "", errors.New("type not support")
	case complex128:
		return "", errors.New("type not support")
	case uintptr:
		return "", errors.New("type not support")
	case string:
		return a.(string), nil
	default: // 其他类型有pointer， struct， array, slice ,map, interface, function, channel
		return "", errors.New("type not support")
	}
}
