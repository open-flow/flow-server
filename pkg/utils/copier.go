package utils

import "github.com/jinzhu/copier"

func DeepCopy(source, target interface{}) error {
	return copier.CopyWithOption(target, source, copier.Option{
		DeepCopy: true,
	})
}
