package utils

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
)

func ConvertMapStructure[T, D any](dto D) (*T, error) {
	entity := new(T)
	config := &mapstructure.DecoderConfig{
		Result:     &entity,
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(dto)
	return entity, err
}

func ConvertDtoToEntity[T, D any](dto D, opts ...copier.Option) (*T, error) {
	entity := new(T)
	err := copier.CopyWithOption(entity, dto, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}
	return entity, nil
}
