// Code generated by protoc-gen-gogo.
// source: github.com/openshift/origin/pkg/project/api/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/openshift/origin/pkg/project/api/v1/generated.proto

	It has these top-level messages:
		Project
		ProjectList
		ProjectRequest
		ProjectSpec
		ProjectStatus
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_kubernetes_pkg_api_v1 "k8s.io/kubernetes/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Project) Reset()                    { *m = Project{} }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ProjectList) Reset()                    { *m = ProjectList{} }
func (*ProjectList) ProtoMessage()               {}
func (*ProjectList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *ProjectRequest) Reset()                    { *m = ProjectRequest{} }
func (*ProjectRequest) ProtoMessage()               {}
func (*ProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *ProjectSpec) Reset()                    { *m = ProjectSpec{} }
func (*ProjectSpec) ProtoMessage()               {}
func (*ProjectSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func (m *ProjectStatus) Reset()                    { *m = ProjectStatus{} }
func (*ProjectStatus) ProtoMessage()               {}
func (*ProjectStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func init() {
	proto.RegisterType((*Project)(nil), "github.com.openshift.origin.pkg.project.api.v1.Project")
	proto.RegisterType((*ProjectList)(nil), "github.com.openshift.origin.pkg.project.api.v1.ProjectList")
	proto.RegisterType((*ProjectRequest)(nil), "github.com.openshift.origin.pkg.project.api.v1.ProjectRequest")
	proto.RegisterType((*ProjectSpec)(nil), "github.com.openshift.origin.pkg.project.api.v1.ProjectSpec")
	proto.RegisterType((*ProjectStatus)(nil), "github.com.openshift.origin.pkg.project.api.v1.ProjectStatus")
}
func (m *Project) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Project) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *ProjectList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProjectList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ProjectRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProjectRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n5, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.DisplayName)))
	i += copy(data[i:], m.DisplayName)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Description)))
	i += copy(data[i:], m.Description)
	return i, nil
}

func (m *ProjectSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProjectSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Finalizers) > 0 {
		for _, s := range m.Finalizers {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *ProjectStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProjectStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Phase)))
	i += copy(data[i:], m.Phase)
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Project) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ProjectList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ProjectRequest) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.DisplayName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Description)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ProjectSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Finalizers) > 0 {
		for _, s := range m.Finalizers {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ProjectStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Phase)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Project) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Project{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ProjectSpec", "ProjectSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ProjectStatus", "ProjectStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_kubernetes_pkg_api_metav1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "Project", "Project", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectRequest{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectSpec{`,
		`Finalizers:` + fmt.Sprintf("%v", this.Finalizers) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProjectStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProjectStatus{`,
		`Phase:` + fmt.Sprintf("%v", this.Phase) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Project) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Project{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalizers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Finalizers = append(m.Finalizers, k8s_io_kubernetes_pkg_api_v1.FinalizerName(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phase = k8s_io_kubernetes_pkg_api_v1.NamespacePhase(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0x4f, 0x6f, 0x12, 0x41,
	0x14, 0x67, 0xa1, 0xd4, 0x32, 0xd8, 0xc6, 0xac, 0x17, 0xc2, 0x61, 0x21, 0xc4, 0x03, 0x51, 0x3b,
	0x13, 0x9a, 0x34, 0x9a, 0x58, 0x3d, 0x6c, 0x8c, 0x89, 0x89, 0x7f, 0x9a, 0xf5, 0x62, 0x4c, 0x7b,
	0x18, 0x96, 0xc7, 0x32, 0x02, 0x3b, 0xe3, 0xcc, 0x2c, 0x89, 0x9e, 0xfc, 0x08, 0x7e, 0x0d, 0xbf,
	0x09, 0xf1, 0xd4, 0x63, 0x4f, 0x44, 0xf0, 0x5b, 0xf4, 0x64, 0x76, 0x76, 0x0a, 0x6b, 0xa9, 0x24,
	0x36, 0xf1, 0xc6, 0x7b, 0xfb, 0xfb, 0xf3, 0xe6, 0xf7, 0x1e, 0xe8, 0x59, 0xc4, 0xf4, 0x20, 0xe9,
	0xe2, 0x90, 0x8f, 0x09, 0x17, 0x10, 0xab, 0x01, 0xeb, 0x6b, 0xc2, 0x25, 0x8b, 0x58, 0x4c, 0xc4,
	0x30, 0x22, 0x42, 0xf2, 0x8f, 0x10, 0x6a, 0x42, 0x05, 0x23, 0x93, 0x0e, 0x89, 0x20, 0x06, 0x49,
	0x35, 0xf4, 0xb0, 0x90, 0x5c, 0x73, 0x17, 0xaf, 0xf8, 0x78, 0xc9, 0xc7, 0x19, 0x1f, 0x8b, 0x61,
	0x84, 0x2d, 0x1f, 0x53, 0xc1, 0xf0, 0xa4, 0x53, 0xdf, 0xcf, 0xf9, 0x45, 0x3c, 0xe2, 0xc4, 0xc8,
	0x74, 0x93, 0xbe, 0xa9, 0x4c, 0x61, 0x7e, 0x65, 0xf2, 0xf5, 0xc3, 0xe1, 0x63, 0x85, 0x19, 0x27,
	0xc3, 0xa4, 0x0b, 0x32, 0x06, 0x0d, 0xca, 0x0c, 0x95, 0x0e, 0x93, 0xc4, 0x13, 0x90, 0x8a, 0xf1,
	0x18, 0x7a, 0x57, 0xa7, 0xaa, 0x3f, 0xfc, 0x3b, 0x6d, 0xfd, 0x0d, 0xf5, 0xfd, 0xeb, 0xd1, 0x32,
	0x89, 0x35, 0x1b, 0xc3, 0x1a, 0xbc, 0x73, 0x3d, 0x3c, 0xd1, 0x6c, 0x44, 0x58, 0xac, 0x95, 0x96,
	0x57, 0x29, 0xad, 0xef, 0x45, 0x74, 0xeb, 0x38, 0x0b, 0xc2, 0x7d, 0x8f, 0x76, 0xc6, 0xa0, 0x69,
	0x8f, 0x6a, 0x5a, 0x73, 0x9a, 0x4e, 0xbb, 0x7a, 0xd0, 0xc6, 0x99, 0x22, 0x5e, 0x29, 0x9a, 0xe8,
	0xb2, 0xc8, 0xf0, 0xdb, 0x6e, 0xca, 0x7b, 0x0d, 0x9a, 0xfa, 0xee, 0x74, 0xd6, 0x28, 0x2c, 0x66,
	0x0d, 0xb4, 0xea, 0x05, 0x4b, 0x35, 0xf7, 0x14, 0x6d, 0x29, 0x01, 0x61, 0xad, 0x68, 0x54, 0x9f,
	0xfc, 0xe3, 0x6a, 0xb0, 0x1d, 0xf0, 0x9d, 0x80, 0xd0, 0xbf, 0x6d, 0x8d, 0xb6, 0xd2, 0x2a, 0x30,
	0xb2, 0x2e, 0xa0, 0x6d, 0xa5, 0xa9, 0x4e, 0x54, 0xad, 0x64, 0x0c, 0x9e, 0xde, 0xd4, 0xc0, 0x88,
	0xf8, 0x7b, 0xd6, 0x62, 0x3b, 0xab, 0x03, 0x2b, 0xde, 0xfa, 0xe1, 0xa0, 0xaa, 0x45, 0xbe, 0x62,
	0x4a, 0xbb, 0xa7, 0x6b, 0x79, 0x91, 0x0d, 0x79, 0xe5, 0xae, 0x02, 0xa7, 0x74, 0x13, 0xdb, 0x1d,
	0x6b, 0xb5, 0x73, 0xd9, 0xc9, 0x85, 0x76, 0x82, 0xca, 0x4c, 0xc3, 0x58, 0xd5, 0x8a, 0xcd, 0x52,
	0xbb, 0x7a, 0xf0, 0xe8, 0x86, 0x8f, 0xf2, 0x77, 0xad, 0x47, 0xf9, 0x65, 0xaa, 0x16, 0x64, 0xa2,
	0xad, 0x73, 0x07, 0xed, 0x59, 0x44, 0x00, 0x9f, 0x12, 0x50, 0xff, 0x73, 0xff, 0x87, 0xa8, 0xda,
	0x63, 0x4a, 0x8c, 0xe8, 0xe7, 0x37, 0x74, 0x0c, 0xe6, 0x0c, 0x2a, 0xfe, 0x5d, 0x4b, 0xa9, 0x3e,
	0x5f, 0x7d, 0x0a, 0xf2, 0x38, 0x43, 0x03, 0x15, 0x4a, 0x26, 0x34, 0xe3, 0xb1, 0x59, 0x6e, 0x9e,
	0xb6, 0xfa, 0x14, 0xe4, 0x71, 0xad, 0xe1, 0x72, 0x4d, 0xe9, 0x8d, 0xb8, 0x27, 0x08, 0xf5, 0x59,
	0x4c, 0x47, 0xec, 0x0b, 0x48, 0x55, 0x73, 0x9a, 0xa5, 0x76, 0xc5, 0x3f, 0x4a, 0x47, 0x7d, 0xb1,
	0xec, 0x5e, 0xcc, 0x1a, 0xf7, 0x37, 0xfd, 0x31, 0xf1, 0x12, 0x6a, 0x86, 0xcc, 0xe9, 0xb5, 0x42,
	0xb4, 0xfb, 0xc7, 0xf5, 0xb8, 0x01, 0x2a, 0x8b, 0x01, 0x55, 0x60, 0x22, 0xac, 0xf8, 0x47, 0x97,
	0xe9, 0x1f, 0xa7, 0xcd, 0x8b, 0x59, 0xe3, 0xc1, 0x46, 0xa3, 0x54, 0x5f, 0x09, 0x1a, 0x82, 0x81,
	0x07, 0x99, 0x94, 0x7f, 0x6f, 0x3a, 0xf7, 0x0a, 0x67, 0x73, 0xaf, 0x70, 0x3e, 0xf7, 0x0a, 0x5f,
	0x17, 0x9e, 0x33, 0x5d, 0x78, 0xce, 0xd9, 0xc2, 0x73, 0x7e, 0x2e, 0x3c, 0xe7, 0xdb, 0x2f, 0xaf,
	0xf0, 0xa1, 0x38, 0xe9, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xb3, 0x68, 0xc8, 0x32, 0x05,
	0x00, 0x00,
}
