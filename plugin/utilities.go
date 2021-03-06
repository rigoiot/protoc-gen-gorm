package plugin

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	gorm "github.com/rigoiot/protoc-gen-gorm/options"
)

func (p *OrmPlugin) getMsgName(o generator.Object) string {
	fqTypeName := p.TypeName(o)
	a := strings.Split(fqTypeName, ".")
	return a[len(a)-1]
}

// retrieves the GormMessageOptions from a message
func getMessageOptions(message *generator.Descriptor) *gorm.GormMessageOptions {
	if message.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(message.Options, gorm.E_Opts)
	if err != nil {
		return nil
	}
	opts, ok := v.(*gorm.GormMessageOptions)
	if !ok {
		return nil
	}
	return opts
}

func getFieldOptions(field *descriptor.FieldDescriptorProto) *gorm.GormFieldOptions {
	if field.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(field.Options, gorm.E_Field)
	if err != nil {
		return nil
	}
	opts, ok := v.(*gorm.GormFieldOptions)
	if !ok {
		return nil
	}
	return opts
}

func getServiceOptions(service *descriptor.ServiceDescriptorProto) *gorm.AutoServerOptions {
	if service.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(service.Options, gorm.E_Server)
	if err != nil {
		return nil
	}
	opts, ok := v.(*gorm.AutoServerOptions)
	if !ok {
		return nil
	}
	return opts
}

func getMethodOptions(method *descriptor.MethodDescriptorProto) *gorm.MethodOptions {
	if method.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(method.Options, gorm.E_Method)
	if err != nil {
		return nil
	}
	opts, ok := v.(*gorm.MethodOptions)
	if !ok {
		return nil
	}
	return opts
}

func isSpecialType(typeName string) bool {
	parts := strings.Split(typeName, ".")
	if len(parts) > 2 { // what kinda format is this????
		panic(fmt.Sprintf(""))
	}
	if len(parts) == 1 { // native to this package = not special
		return false
	}
	// anything that looks like a google_protobufX should be considered special
	if strings.HasPrefix(strings.TrimLeft(typeName, "[]*"), "google_protobuf") {
		return true
	}
	switch parts[len(parts)-1] {
	case protoTypeJSON,
		protoTypeUUID,
		protoTypeUUIDValue,
		protoTypeResource,
		protoTypeInet:
		return true
	}
	return false
}

// allFiles return the all FileDescriptor, hack teh allFiles in Generator
func (p *OrmPlugin) allFiles() []*generator.FileDescriptor {

	v := reflect.Indirect(reflect.ValueOf(p.Generator))
	m := v.FieldByName("allFiles")
	pv := unsafe.Pointer(m.UnsafeAddr())

	return *(*[]*generator.FileDescriptor)(pv)
}

// if p.useStatusError {
// 	p.P(`return nil, `, p.Import(statusImport), `.Error(`, p.Import(codesImport), `.`, `"Database Transaction For Request Missing")`)
// } else {
// 	p.UsingGoImports("errors")
// 	p.P(`return nil, errors.New("Database Transaction For Request Missing")`)
// }
func (p *OrmPlugin) newError(code string) string {
	if p.useStatusError {
		if len(code) == 0 {
			code = "Unknown"
		}
		return p.Import(statusImport) + ".Error(" + p.Import(codesImport) + "." + code + ", "
	} else {
		p.UsingGoImports("errors")
		return "errors.New("
	}
}
