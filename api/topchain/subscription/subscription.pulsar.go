// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package subscription

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Subscription             protoreflect.MessageDescriptor
	fd_Subscription_id          protoreflect.FieldDescriptor
	fd_Subscription_deal_id     protoreflect.FieldDescriptor
	fd_Subscription_provider    protoreflect.FieldDescriptor
	fd_Subscription_start_block protoreflect.FieldDescriptor
	fd_Subscription_end_block   protoreflect.FieldDescriptor
)

func init() {
	file_topchain_subscription_subscription_proto_init()
	md_Subscription = File_topchain_subscription_subscription_proto.Messages().ByName("Subscription")
	fd_Subscription_id = md_Subscription.Fields().ByName("id")
	fd_Subscription_deal_id = md_Subscription.Fields().ByName("deal_id")
	fd_Subscription_provider = md_Subscription.Fields().ByName("provider")
	fd_Subscription_start_block = md_Subscription.Fields().ByName("start_block")
	fd_Subscription_end_block = md_Subscription.Fields().ByName("end_block")
}

var _ protoreflect.Message = (*fastReflection_Subscription)(nil)

type fastReflection_Subscription Subscription

func (x *Subscription) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Subscription)(x)
}

func (x *Subscription) slowProtoReflect() protoreflect.Message {
	mi := &file_topchain_subscription_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Subscription_messageType fastReflection_Subscription_messageType
var _ protoreflect.MessageType = fastReflection_Subscription_messageType{}

type fastReflection_Subscription_messageType struct{}

func (x fastReflection_Subscription_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Subscription)(nil)
}
func (x fastReflection_Subscription_messageType) New() protoreflect.Message {
	return new(fastReflection_Subscription)
}
func (x fastReflection_Subscription_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Subscription
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Subscription) Descriptor() protoreflect.MessageDescriptor {
	return md_Subscription
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Subscription) Type() protoreflect.MessageType {
	return _fastReflection_Subscription_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Subscription) New() protoreflect.Message {
	return new(fastReflection_Subscription)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Subscription) Interface() protoreflect.ProtoMessage {
	return (*Subscription)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Subscription) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != "" {
		value := protoreflect.ValueOfString(x.Id)
		if !f(fd_Subscription_id, value) {
			return
		}
	}
	if x.DealId != "" {
		value := protoreflect.ValueOfString(x.DealId)
		if !f(fd_Subscription_deal_id, value) {
			return
		}
	}
	if x.Provider != "" {
		value := protoreflect.ValueOfString(x.Provider)
		if !f(fd_Subscription_provider, value) {
			return
		}
	}
	if x.StartBlock != uint64(0) {
		value := protoreflect.ValueOfUint64(x.StartBlock)
		if !f(fd_Subscription_start_block, value) {
			return
		}
	}
	if x.EndBlock != uint64(0) {
		value := protoreflect.ValueOfUint64(x.EndBlock)
		if !f(fd_Subscription_end_block, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Subscription) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "topchain.subscription.Subscription.id":
		return x.Id != ""
	case "topchain.subscription.Subscription.deal_id":
		return x.DealId != ""
	case "topchain.subscription.Subscription.provider":
		return x.Provider != ""
	case "topchain.subscription.Subscription.start_block":
		return x.StartBlock != uint64(0)
	case "topchain.subscription.Subscription.end_block":
		return x.EndBlock != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Subscription) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "topchain.subscription.Subscription.id":
		x.Id = ""
	case "topchain.subscription.Subscription.deal_id":
		x.DealId = ""
	case "topchain.subscription.Subscription.provider":
		x.Provider = ""
	case "topchain.subscription.Subscription.start_block":
		x.StartBlock = uint64(0)
	case "topchain.subscription.Subscription.end_block":
		x.EndBlock = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Subscription) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "topchain.subscription.Subscription.id":
		value := x.Id
		return protoreflect.ValueOfString(value)
	case "topchain.subscription.Subscription.deal_id":
		value := x.DealId
		return protoreflect.ValueOfString(value)
	case "topchain.subscription.Subscription.provider":
		value := x.Provider
		return protoreflect.ValueOfString(value)
	case "topchain.subscription.Subscription.start_block":
		value := x.StartBlock
		return protoreflect.ValueOfUint64(value)
	case "topchain.subscription.Subscription.end_block":
		value := x.EndBlock
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Subscription) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "topchain.subscription.Subscription.id":
		x.Id = value.Interface().(string)
	case "topchain.subscription.Subscription.deal_id":
		x.DealId = value.Interface().(string)
	case "topchain.subscription.Subscription.provider":
		x.Provider = value.Interface().(string)
	case "topchain.subscription.Subscription.start_block":
		x.StartBlock = value.Uint()
	case "topchain.subscription.Subscription.end_block":
		x.EndBlock = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Subscription) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "topchain.subscription.Subscription.id":
		panic(fmt.Errorf("field id of message topchain.subscription.Subscription is not mutable"))
	case "topchain.subscription.Subscription.deal_id":
		panic(fmt.Errorf("field deal_id of message topchain.subscription.Subscription is not mutable"))
	case "topchain.subscription.Subscription.provider":
		panic(fmt.Errorf("field provider of message topchain.subscription.Subscription is not mutable"))
	case "topchain.subscription.Subscription.start_block":
		panic(fmt.Errorf("field start_block of message topchain.subscription.Subscription is not mutable"))
	case "topchain.subscription.Subscription.end_block":
		panic(fmt.Errorf("field end_block of message topchain.subscription.Subscription is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Subscription) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "topchain.subscription.Subscription.id":
		return protoreflect.ValueOfString("")
	case "topchain.subscription.Subscription.deal_id":
		return protoreflect.ValueOfString("")
	case "topchain.subscription.Subscription.provider":
		return protoreflect.ValueOfString("")
	case "topchain.subscription.Subscription.start_block":
		return protoreflect.ValueOfUint64(uint64(0))
	case "topchain.subscription.Subscription.end_block":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: topchain.subscription.Subscription"))
		}
		panic(fmt.Errorf("message topchain.subscription.Subscription does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Subscription) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in topchain.subscription.Subscription", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Subscription) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Subscription) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Subscription) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Subscription) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Subscription)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Id)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DealId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Provider)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartBlock != 0 {
			n += 1 + runtime.Sov(uint64(x.StartBlock))
		}
		if x.EndBlock != 0 {
			n += 1 + runtime.Sov(uint64(x.EndBlock))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Subscription)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.EndBlock != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndBlock))
			i--
			dAtA[i] = 0x28
		}
		if x.StartBlock != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartBlock))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Provider) > 0 {
			i -= len(x.Provider)
			copy(dAtA[i:], x.Provider)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Provider)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.DealId) > 0 {
			i -= len(x.DealId)
			copy(dAtA[i:], x.DealId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DealId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Id) > 0 {
			i -= len(x.Id)
			copy(dAtA[i:], x.Id)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Id)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Subscription)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Subscription: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Subscription: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Id = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DealId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Provider = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
				}
				x.StartBlock = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartBlock |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
				}
				x.EndBlock = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndBlock |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: topchain/subscription/subscription.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DealId     string `protobuf:"bytes,2,opt,name=deal_id,json=dealId,proto3" json:"deal_id,omitempty"`
	Provider   string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	StartBlock uint64 `protobuf:"varint,4,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock   uint64 `protobuf:"varint,5,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topchain_subscription_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_topchain_subscription_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription) GetDealId() string {
	if x != nil {
		return x.DealId
	}
	return ""
}

func (x *Subscription) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Subscription) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *Subscription) GetEndBlock() uint64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

var File_topchain_subscription_subscription_proto protoreflect.FileDescriptor

var file_topchain_subscription_subscription_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x6f, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x6f, 0x70, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0xc7, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f,
	0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x74, 0x6f, 0x70, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x54,
	0x53, 0x58, 0xaa, 0x02, 0x15, 0x54, 0x6f, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x15, 0x54, 0x6f, 0x70,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0xe2, 0x02, 0x21, 0x54, 0x6f, 0x70, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x54, 0x6f, 0x70, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x3a, 0x3a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_topchain_subscription_subscription_proto_rawDescOnce sync.Once
	file_topchain_subscription_subscription_proto_rawDescData = file_topchain_subscription_subscription_proto_rawDesc
)

func file_topchain_subscription_subscription_proto_rawDescGZIP() []byte {
	file_topchain_subscription_subscription_proto_rawDescOnce.Do(func() {
		file_topchain_subscription_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_topchain_subscription_subscription_proto_rawDescData)
	})
	return file_topchain_subscription_subscription_proto_rawDescData
}

var file_topchain_subscription_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_topchain_subscription_subscription_proto_goTypes = []interface{}{
	(*Subscription)(nil), // 0: topchain.subscription.Subscription
}
var file_topchain_subscription_subscription_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_topchain_subscription_subscription_proto_init() }
func file_topchain_subscription_subscription_proto_init() {
	if File_topchain_subscription_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topchain_subscription_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_topchain_subscription_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_topchain_subscription_subscription_proto_goTypes,
		DependencyIndexes: file_topchain_subscription_subscription_proto_depIdxs,
		MessageInfos:      file_topchain_subscription_subscription_proto_msgTypes,
	}.Build()
	File_topchain_subscription_subscription_proto = out.File
	file_topchain_subscription_subscription_proto_rawDesc = nil
	file_topchain_subscription_subscription_proto_goTypes = nil
	file_topchain_subscription_subscription_proto_depIdxs = nil
}
