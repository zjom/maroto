// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mappers "github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	mock "github.com/stretchr/testify/mock"
)

// AbstractFactoryMaps is an autogenerated mock type for the AbstractFactoryMaps type
type AbstractFactoryMaps struct {
	mock.Mock
}

type AbstractFactoryMaps_Expecter struct {
	mock *mock.Mock
}

func (_m *AbstractFactoryMaps) EXPECT() *AbstractFactoryMaps_Expecter {
	return &AbstractFactoryMaps_Expecter{mock: &_m.Mock}
}

// NewBarcode provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewBarcode(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewBarcode")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewBarcode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewBarcode'
type AbstractFactoryMaps_NewBarcode_Call struct {
	*mock.Call
}

// NewBarcode is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewBarcode(document interface{}) *AbstractFactoryMaps_NewBarcode_Call {
	return &AbstractFactoryMaps_NewBarcode_Call{Call: _e.mock.On("NewBarcode", document)}
}

func (_c *AbstractFactoryMaps_NewBarcode_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewBarcode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewBarcode_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewBarcode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewBarcode_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewBarcode_Call {
	_c.Call.Return(run)
	return _c
}

// NewCol provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewCol(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewCol")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewCol_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewCol'
type AbstractFactoryMaps_NewCol_Call struct {
	*mock.Call
}

// NewCol is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewCol(document interface{}) *AbstractFactoryMaps_NewCol_Call {
	return &AbstractFactoryMaps_NewCol_Call{Call: _e.mock.On("NewCol", document)}
}

func (_c *AbstractFactoryMaps_NewCol_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewCol_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewCol_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewCol_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewCol_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewCol_Call {
	_c.Call.Return(run)
	return _c
}

// NewImage provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewImage(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewImage")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewImage'
type AbstractFactoryMaps_NewImage_Call struct {
	*mock.Call
}

// NewImage is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewImage(document interface{}) *AbstractFactoryMaps_NewImage_Call {
	return &AbstractFactoryMaps_NewImage_Call{Call: _e.mock.On("NewImage", document)}
}

func (_c *AbstractFactoryMaps_NewImage_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewImage_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewImage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewImage_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewImage_Call {
	_c.Call.Return(run)
	return _c
}

// NewLine provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewLine(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewLine")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewLine'
type AbstractFactoryMaps_NewLine_Call struct {
	*mock.Call
}

// NewLine is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewLine(document interface{}) *AbstractFactoryMaps_NewLine_Call {
	return &AbstractFactoryMaps_NewLine_Call{Call: _e.mock.On("NewLine", document)}
}

func (_c *AbstractFactoryMaps_NewLine_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewLine_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewLine_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewLine_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewLine_Call {
	_c.Call.Return(run)
	return _c
}

// NewList provides a mock function with given fields: document, sourceKey, generate
func (_m *AbstractFactoryMaps) NewList(document interface{}, sourceKey string, generate mappers.GenerateComponent) (mappers.Componentmapper, error) {
	ret := _m.Called(document, sourceKey, generate)

	if len(ret) == 0 {
		panic("no return value specified for NewList")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, string, mappers.GenerateComponent) (mappers.Componentmapper, error)); ok {
		return rf(document, sourceKey, generate)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, mappers.GenerateComponent) mappers.Componentmapper); ok {
		r0 = rf(document, sourceKey, generate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, mappers.GenerateComponent) error); ok {
		r1 = rf(document, sourceKey, generate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewList'
type AbstractFactoryMaps_NewList_Call struct {
	*mock.Call
}

// NewList is a helper method to define mock.On call
//   - document interface{}
//   - sourceKey string
//   - generate mappers.GenerateComponent
func (_e *AbstractFactoryMaps_Expecter) NewList(document interface{}, sourceKey interface{}, generate interface{}) *AbstractFactoryMaps_NewList_Call {
	return &AbstractFactoryMaps_NewList_Call{Call: _e.mock.On("NewList", document, sourceKey, generate)}
}

func (_c *AbstractFactoryMaps_NewList_Call) Run(run func(document interface{}, sourceKey string, generate mappers.GenerateComponent)) *AbstractFactoryMaps_NewList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(string), args[2].(mappers.GenerateComponent))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewList_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewList_Call) RunAndReturn(run func(interface{}, string, mappers.GenerateComponent) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMatrixcode provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewMatrixcode(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewMatrixcode")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewMatrixcode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewMatrixcode'
type AbstractFactoryMaps_NewMatrixcode_Call struct {
	*mock.Call
}

// NewMatrixcode is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewMatrixcode(document interface{}) *AbstractFactoryMaps_NewMatrixcode_Call {
	return &AbstractFactoryMaps_NewMatrixcode_Call{Call: _e.mock.On("NewMatrixcode", document)}
}

func (_c *AbstractFactoryMaps_NewMatrixcode_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewMatrixcode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewMatrixcode_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewMatrixcode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewMatrixcode_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewMatrixcode_Call {
	_c.Call.Return(run)
	return _c
}

// NewPage provides a mock function with given fields: document, sourceKey
func (_m *AbstractFactoryMaps) NewPage(document interface{}, sourceKey string) (mappers.Componentmapper, error) {
	ret := _m.Called(document, sourceKey)

	if len(ret) == 0 {
		panic("no return value specified for NewPage")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, string) (mappers.Componentmapper, error)); ok {
		return rf(document, sourceKey)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string) mappers.Componentmapper); ok {
		r0 = rf(document, sourceKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(document, sourceKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPage'
type AbstractFactoryMaps_NewPage_Call struct {
	*mock.Call
}

// NewPage is a helper method to define mock.On call
//   - document interface{}
//   - sourceKey string
func (_e *AbstractFactoryMaps_Expecter) NewPage(document interface{}, sourceKey interface{}) *AbstractFactoryMaps_NewPage_Call {
	return &AbstractFactoryMaps_NewPage_Call{Call: _e.mock.On("NewPage", document, sourceKey)}
}

func (_c *AbstractFactoryMaps_NewPage_Call) Run(run func(document interface{}, sourceKey string)) *AbstractFactoryMaps_NewPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(string))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewPage_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewPage_Call) RunAndReturn(run func(interface{}, string) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewPage_Call {
	_c.Call.Return(run)
	return _c
}

// NewQrcode provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewQrcode(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewQrcode")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewQrcode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewQrcode'
type AbstractFactoryMaps_NewQrcode_Call struct {
	*mock.Call
}

// NewQrcode is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewQrcode(document interface{}) *AbstractFactoryMaps_NewQrcode_Call {
	return &AbstractFactoryMaps_NewQrcode_Call{Call: _e.mock.On("NewQrcode", document)}
}

func (_c *AbstractFactoryMaps_NewQrcode_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewQrcode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewQrcode_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewQrcode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewQrcode_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewQrcode_Call {
	_c.Call.Return(run)
	return _c
}

// NewRow provides a mock function with given fields: document, sourceKey
func (_m *AbstractFactoryMaps) NewRow(document interface{}, sourceKey string) (mappers.Componentmapper, error) {
	ret := _m.Called(document, sourceKey)

	if len(ret) == 0 {
		panic("no return value specified for NewRow")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, string) (mappers.Componentmapper, error)); ok {
		return rf(document, sourceKey)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string) mappers.Componentmapper); ok {
		r0 = rf(document, sourceKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(document, sourceKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewRow'
type AbstractFactoryMaps_NewRow_Call struct {
	*mock.Call
}

// NewRow is a helper method to define mock.On call
//   - document interface{}
//   - sourceKey string
func (_e *AbstractFactoryMaps_Expecter) NewRow(document interface{}, sourceKey interface{}) *AbstractFactoryMaps_NewRow_Call {
	return &AbstractFactoryMaps_NewRow_Call{Call: _e.mock.On("NewRow", document, sourceKey)}
}

func (_c *AbstractFactoryMaps_NewRow_Call) Run(run func(document interface{}, sourceKey string)) *AbstractFactoryMaps_NewRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(string))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewRow_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewRow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewRow_Call) RunAndReturn(run func(interface{}, string) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewSignature provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewSignature(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewSignature")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewSignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewSignature'
type AbstractFactoryMaps_NewSignature_Call struct {
	*mock.Call
}

// NewSignature is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewSignature(document interface{}) *AbstractFactoryMaps_NewSignature_Call {
	return &AbstractFactoryMaps_NewSignature_Call{Call: _e.mock.On("NewSignature", document)}
}

func (_c *AbstractFactoryMaps_NewSignature_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewSignature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewSignature_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewSignature_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewSignature_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewSignature_Call {
	_c.Call.Return(run)
	return _c
}

// NewText provides a mock function with given fields: document
func (_m *AbstractFactoryMaps) NewText(document interface{}) (mappers.Componentmapper, error) {
	ret := _m.Called(document)

	if len(ret) == 0 {
		panic("no return value specified for NewText")
	}

	var r0 mappers.Componentmapper
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (mappers.Componentmapper, error)); ok {
		return rf(document)
	}
	if rf, ok := ret.Get(0).(func(interface{}) mappers.Componentmapper); ok {
		r0 = rf(document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.Componentmapper)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractFactoryMaps_NewText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewText'
type AbstractFactoryMaps_NewText_Call struct {
	*mock.Call
}

// NewText is a helper method to define mock.On call
//   - document interface{}
func (_e *AbstractFactoryMaps_Expecter) NewText(document interface{}) *AbstractFactoryMaps_NewText_Call {
	return &AbstractFactoryMaps_NewText_Call{Call: _e.mock.On("NewText", document)}
}

func (_c *AbstractFactoryMaps_NewText_Call) Run(run func(document interface{})) *AbstractFactoryMaps_NewText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AbstractFactoryMaps_NewText_Call) Return(_a0 mappers.Componentmapper, _a1 error) *AbstractFactoryMaps_NewText_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AbstractFactoryMaps_NewText_Call) RunAndReturn(run func(interface{}) (mappers.Componentmapper, error)) *AbstractFactoryMaps_NewText_Call {
	_c.Call.Return(run)
	return _c
}

// NewAbstractFactoryMaps creates a new instance of AbstractFactoryMaps. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAbstractFactoryMaps(t interface {
	mock.TestingT
	Cleanup(func())
},
) *AbstractFactoryMaps {
	mock := &AbstractFactoryMaps{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}