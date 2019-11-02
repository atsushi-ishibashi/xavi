package xavi

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	tagName string = "xavi"
)

func Pass(dst, src interface{}) error {
	if err := validateDst(dst); err != nil {
		return err
	}
	if err := validateSrc(src); err != nil {
		return err
	}
	dstVo := reflect.ValueOf(dst).Elem()
	srcVo := reflect.ValueOf(src)
	switch srcVo.Kind() {
	case reflect.Ptr, reflect.Interface:
		srcVo = srcVo.Elem()
	}
	srcType := srcVo.Type()
	dstType := dstVo.Type()
	for i := 0; i < srcType.NumField(); i++ {
		srcF := srcType.Field(i)
		srcV := srcVo.Field(i)
		for ii := 0; ii < dstType.NumField(); ii++ {
			dstF := dstType.Field(ii)
			if matchTag(dstF, srcF) || matchKeyKind(dstF, srcF) {
				dstVo.Field(ii).Set(srcV)
				break
			}
		}
	}
	return nil
}

func matchTag(dstField, srcField reflect.StructField) bool {
	srcTag := srcField.Tag.Get(tagName)
	if srcTag == "" {
		return false
	}
	dstTag := dstField.Tag.Get(tagName)
	if dstTag == "" {
		return false
	}
	if srcTag != dstTag {
		return false
	}
	return matchKind(dstField, srcField)
}

func matchKeyKind(dstField, srcField reflect.StructField) bool {
	return srcField.Name == dstField.Name && matchKind(dstField, srcField)
}

func matchKind(dstField, srcField reflect.StructField) bool {
	if srcField.Type.Kind() != dstField.Type.Kind() {
		return false
	}
	srcKind := srcField.Type.Kind()
	switch srcKind {
	case reflect.Struct, reflect.Ptr, reflect.Slice, reflect.Array, reflect.Func, reflect.Map:
		return srcField.Type.Name() == dstField.Type.Name()
	case reflect.Interface:
		//FIXME
		return srcField.Type.Name() == dstField.Type.Name()
	}
	return true
}

func validateSrc(src interface{}) error {
	vo := reflect.ValueOf(src)
	switch vo.Kind() {
	case reflect.Ptr, reflect.Interface:
		if vo.IsNil() {
			return errors.New("src is nil")
		}
	case reflect.Struct:
	default:
		return fmt.Errorf("unexpected src kind: %s", vo.Kind().String())
	}
	return nil
}

func validateDst(dst interface{}) error {
	vo := reflect.ValueOf(dst)
	if vo.Kind() != reflect.Ptr {
		return fmt.Errorf("expected dst kind=%s, got %s", reflect.Ptr.String(), vo.Kind().String())
	}
	if vo.IsNil() {
		return errors.New("dst is nil")
	}
	if elm := vo.Elem(); elm.Kind() != reflect.Struct {
		return fmt.Errorf("expected dst element kind=%s, got %s", reflect.Struct.String(), elm.Kind().String())
	}
	return nil
}
