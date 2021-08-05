package utils

import "github.com/jinzhu/copier"

func DeepCopy(target interface{}, source interface{}) error {
	return copier.CopyWithOption(target, source, copier.Option{
		DeepCopy: true,
	})
}
