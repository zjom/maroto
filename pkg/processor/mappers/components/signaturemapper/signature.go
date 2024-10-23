// Package signature implements creation of signatures.
package signaturemapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Signature struct {
	SourceKey string
	Value     string
	Props     *propsmapper.Signature
}

func NewSignature(code interface{}) (*Signature, error) {
	signatureMap, ok := code.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure signature can be converted to map[string] interface{}")
	}

	signature := &Signature{}
	if err := signature.addFields(signatureMap); err != nil {
		return nil, err
	}
	if err := signature.validateFields(); err != nil {
		return nil, err
	}

	return signature, nil
}

// addFields is responsible for adding the signature fields according to
// the properties informed in the map
func (s *Signature) addFields(csignatureMap map[string]interface{}) error {
	fieldMappers := s.getFieldMappers()

	for templateName, template := range csignatureMap {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the signature cannot be mapped to any valid field", templateName)
		}
		err := mapper(template)
		if err != nil {
			return err
		}
	}
	return nil
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (s *Signature) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"source_key": s.setSourceKey,
		"value":      s.setValue,
		"props":      s.setProps,
	}
}

func (s *Signature) setSourceKey(template interface{}) error {
	sourceKey, ok := template.(string)
	if !ok {
		return fmt.Errorf("source_key cannot be converted to a string")
	}
	s.SourceKey = sourceKey
	return nil
}

func (s *Signature) setValue(template interface{}) error {
	value, ok := template.(string)
	if !ok {
		return fmt.Errorf("value cannot be converted to a string")
	}
	s.Value = value
	return nil
}

func (s *Signature) setProps(template interface{}) error {
	props, err := propsmapper.NewSignature(template)
	if err != nil {
		return err
	}
	s.Props = props
	return nil
}

func (s *Signature) validateFields() error {
	if s.Value == "" && s.SourceKey == "" {
		return fmt.Errorf("no value passed for signature. Add the 'source_key' or a value")
	}
	return nil
}

func (b *Signature) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
