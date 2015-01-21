// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package builtin provides documentation for Go's predeclared identifiers.
	The items documented here are not actually in package builtin
	but their descriptions here allow godoc to present documentation
	for the language's special identifiers.
*/

/*
	builtin ��ΪGo��Ԥ������ʶ���ṩ���ĵ�.
	�˴��г�����Ŀ��ʵ������ buildin ���У������ǵ�����ֻ��Ϊ���� godoc
	�������Ե������ʶ���ṩ�ĵ���
*/
package builtin

// bool is the set of boolean values, true and false.

// bool �ǲ���ֵ�ļ��ϣ��� true �� false��
type bool bool

// true and false are the two untyped boolean values.

// true �� false �����������Ͳ���ֵ��
const (
	// �����Ͳ�����
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)

// uint8 is the set of all unsigned 8-bit integers.
// Range: 0 through 255.

// uint8 �������޷���8λ�����ļ��ϡ�
// ��Χ��0 �� 255��
type uint8 uint8

// uint16 is the set of all unsigned 16-bit integers.
// Range: 0 through 65535.

// uint16 �������޷���16λ�����ļ��ϡ�
// ��Χ��0 �� 65535��
type uint16 uint16

// uint32 is the set of all unsigned 32-bit integers.
// Range: 0 through 4294967295.

// uint32 �������޷���32λ�����ļ��ϡ�
// ��Χ��0 �� 4294967295��
type uint32 uint32

// uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.

// uint64 �������޷���64λ�����ļ��ϡ�
// ��Χ��0 �� 18446744073709551615��
type uint64 uint64

// int8 is the set of all signed 8-bit integers.
// Range: -128 through 127.

// int8 �����д�����8λ�����ļ��ϡ�
// ��Χ��-128 �� 127��
type int8 int8

// int16 is the set of all signed 16-bit integers.
// Range: -32768 through 32767.

// int16 �����д�����16λ�����ļ��ϡ�
// ��Χ��-32768 �� 32767��
type int16 int16

// int32 is the set of all signed 32-bit integers.
// Range: -2147483648 through 2147483647.

// int32 �����д�����32λ�����ļ��ϡ�
// ��Χ��-2147483648 �� 2147483647��
type int32 int32

// int64 is the set of all signed 64-bit integers.
// Range: -9223372036854775808 through 9223372036854775807.

// int64 �����д�����64λ�����ļ��ϡ�
// ��Χ��-9223372036854775808 �� 9223372036854775807��
type int64 int64

// float32 is the set of all IEEE-754 32-bit floating-point numbers.

// float32 ������IEEE-754 32λ�������ļ��ϡ�
type float32 float32

// float64 is the set of all IEEE-754 64-bit floating-point numbers.

// float64 ������IEEE-754 64λ�������ļ��ϡ�
type float64 float64

// complex64 is the set of all complex numbers with float32 real and
// imaginary parts.

// complex64 ������ʵ�����鲿Ϊ float32 �ĸ������ϡ�
type complex64 complex64

// complex128 is the set of all complex numbers with float64 real and
// imaginary parts.

// complex128 ������ʵ�����鲿Ϊ float64 �ĸ������ϡ�
type complex128 complex128

// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
// not nil. Values of string type are immutable.

// string ������8λ�ֽڵ��ַ������ϣ�ϰ�������ڴ�����UTF-8������ı���������������ˡ�
// string ��Ϊ�գ�����Ϊ nil��string ���͵�ֵ�ǲ���ġ�
type string string

// int is a signed integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, int32.

// int �Ǵ������������ͣ����С����Ϊ32λ��
// ����һ��ȷ�е����ͣ������� int32 �ı�����
type int int

// uint is an unsigned integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, uint32.

// uint ���޷����������ͣ����С����Ϊ32λ��
// ����һ��ȷ�е����ͣ������� uint32 �ı�����
type uint uint

// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.

// uintptr Ϊ�������ͣ����С���������κ�ָ���λģʽ��
type uintptr uintptr

// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.

// byte Ϊ uint8 �ı���������ȫ�ȼ��� uint8��
// ϰ���������������ֽ�ֵ��8λ�޷�������ֵ��
type byte byte

// rune is an alias for int32 and is equivalent int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.

// rune Ϊ int32 �ı���������ȫ�ȼ��� int32��
// ϰ���������������ַ�ֵ������ֵ��
type rune rune

// iota is a predeclared identifier representing the untyped integer ordinal
// number of the current const specification in a (usually parenthesized)
// const declaration. It is zero-indexed.

// iota ΪԤ�����ı�ʶ��������ʾ���������У�һ���������У���
// ��ǰ�����淶�������ͻ���������������0��ʼ������
const iota = 0 // Untyped int. // �����ͻ� int��

// nil is a predeclared identifier representing the zero value for a
// pointer, channel, func, interface, map, or slice type.

// nil ΪԤ�����ı�ʾ��������ʾָ�롢�ŵ����������ӿڡ�ӳ�����Ƭ���͵���ֵ��
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
// Type ����Ϊָ�롢�ŵ����������ӿڡ�ӳ�����Ƭ���͡�

// Type is here for the purposes of documentation only. It is a stand-in
// for any Go type, but represents the same type for any given function
// invocation.

// Type �ڴ�ֻ�����ĵ�Ŀ�ġ�
// ����������Go�����ͣ��������κθ����ĺ���������˵����������������ͬ�����͡�
type Type int

// Type1 is here for the purposes of documentation only. It is a stand-in
// for any Go type, but represents the same type for any given function
// invocation.

// Type1 �ڴ�ֻ�����ĵ�Ŀ�ġ�
// ����������Go�����ͣ��������κθ����ĺ���������˵����������������ͬ�����͡�
type Type1 int

// IntegerType is here for the purposes of documentation only. It is a stand-in
// for any integer type: int, uint, int8 etc.

// IntegerType �ڴ�ֻ�����ĵ�Ŀ�ġ�
// ���������е��������ͣ��� int��uint��int8 �ȡ�
type IntegerType int

// FloatType is here for the purposes of documentation only. It is a stand-in
// for either float type: float32 or float64.

// FloatType �ڴ�ֻ�����ĵ�Ŀ�ġ�
// ���������еĸ��������ͣ��� float32 �� float64��
type FloatType float32

// ComplexType is here for the purposes of documentation only. It is a
// stand-in for either complex type: complex64 or complex128.

// ComplexType �ڴ�ֻ�����ĵ�Ŀ�ġ�
// ���������еĸ������ͣ��� complex64 �� complex128��
type ComplexType complex64

// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:
//	slice = append([]byte("hello "), "world"...)

// append �ڽ�������Ԫ��׷�ӵ���Ƭ��ĩβ��
// �������㹻����������Ŀ��ͻ�������Ƭ�������µ�Ԫ�ء����򣬾ͻ����һ���µĻ������顣
// append ���ظ��º����Ƭ����˱���洢׷�Ӻ�Ľ����ͨ��Ϊ��������Ƭ�����ı�����
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
// ��Ϊһ���������������ַ�׷�ӵ��ֽ�����֮���ǺϷ��ģ�����������
//	slice = append([]byte("hello "), "world"...)
func append(slice []Type, elems ...Type) []Type

// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).

// copy �ڽ�������Ԫ�ش���Դ��Ƭ���Ƶ�Ŀ����Ƭ�С�
// ����������ǣ���Ҳ�ܽ��ֽڴ��ַ������Ƶ��ֽ���Ƭ�У�����Դ��Ŀ������ص���
// copy ���ر����Ƶ�Ԫ�������������� len(src) �� len(dst) �н�С���Ǹ���
func copy(dst, src []Type) int

// The delete built-in function deletes the element with the specified key
// (m[key]) from the map. If m is nil or there is no such element, delete
// is a no-op.

// delete �ڽ���������ָ���ļ���Ԫ�ش�ӳ����ɾ����
// �� m Ϊ nil ���޴�Ԫ�أ�delete ��Ϊ�ղ�����
func delete(m map[Type]Type1, key Type)

// The len built-in function returns the length of v, according to its type:
//	Array: the number of elements in v.
//	Pointer to array: the number of elements in *v (even if v is nil).
//	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
//	String: the number of bytes in v.
//	Channel: the number of elements queued (unread) in the channel buffer;
//	if v is nil, len(v) is zero.

// len �ڽ��������� v �ĳ��ȣ���ȡ���ھ������ͣ�
//	���飺v ��Ԫ�ص�������
//	����ָ�룺*v ��Ԫ�ص���������ʹ v Ϊ nil����
//	��Ƭ��ӳ�䣺v ��Ԫ�ص��������� v Ϊ nil��len(v) ��Ϊ�㡣
//	�ַ�����v ���ֽڵ�������
//	�ŵ����ŵ������ж��У�δ��ȡ��Ԫ�ص��������� v Ϊ nil��len(v) ��Ϊ�㡣
func len(v Type) int

// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.

// cap �ڽ��������� v ����������ȡ���ھ������ͣ�
//	���飺v ��Ԫ�ص��������� len(v) ��ͬ����
//	����ָ�룺*v ��Ԫ�ص��������� len(v) ��ͬ����
//	��Ƭ����������Ƭʱ����Ƭ�ܹ��ﵽ����󳤶ȣ��� v Ϊ nil��len(v) ��Ϊ�㡣
//	�ŵ�������Ԫ�صĵ�Ԫ����Ӧ�ŵ�������������� v Ϊ nil��len(v) ��Ϊ�㡣
func cap(v Type) int

// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length, so make([]int, 0, 10) allocates a slice of length 0 and
//	capacity 10.
//	Map: An initial allocation is made according to the size but the
//	resulting map has length 0. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.

// make �ڽ��������䲢��ʼ��һ������Ϊ��Ƭ��ӳ�䡢�򣨽���Ϊ���ŵ��Ķ���
// �� new ��ͬ���ǣ����һ��ʵ��Ϊ���ͣ�����ֵ����ͬ���ǣ�make �ķ�������
// ���������ͬ������ָ������ָ�롣�������ȡ���ھ�������ͣ�
//	��Ƭ��size ָ�����䳤�ȡ�����Ƭ�����������䳤�ȡ��ڶ�������ʵ�ο�����ָ��
//		��ͬ�������������벻С���䳤�ȣ���� make([]int, 0, 10) �����һ������Ϊ0��
//		����Ϊ10����Ƭ��
//	ӳ�䣺��ʼ����Ĵ���ȡ���� size����������ӳ�䳤��Ϊ0��size ����ʡ�ԣ����������
//		�ͻ����һ��С����ʼ��С��
//	�ŵ����ŵ��Ļ������ָ���Ļ���������ʼ������ size Ϊ���ʡ�ԣ����ŵ���Ϊ�޻���ġ�
func make(Type, size IntegerType) Type

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.

// new �ڽ����������ڴ档
// ���һ��ʵ��Ϊ���ͣ�����ֵ���䷵��ֵΪָ������͵��·������ֵ��ָ�롣
func new(Type) *Type

// The complex built-in function constructs a complex value from two
// floating-point values. The real and imaginary parts must be of the same
// size, either float32 or float64 (or assignable to them), and the return
// value will be the corresponding complex type (complex64 for float32,
// complex128 for float64).

// complex �ڽ�����������������ֵ�����һ������ֵ��
// ��ʵ�����鲿�Ĵ�С������ͬ���� float32 �� float64����ɸ������ǵģ����䷵��ֵ
// ��Ϊ��Ӧ�ĸ������ͣ�complex64 ��Ӧ float32��complex128 ��Ӧ float64����
func complex(r, i FloatType) ComplexType

// The real built-in function returns the real part of the complex number c.
// The return value will be floating point type corresponding to the type of c.

// real �ڽ��������ظ��� c ��ʵ����
// �䷵��ֵΪ��Ӧ�� c ���͵ĸ�������
func real(c ComplexType) FloatType

// The imag built-in function returns the imaginary part of the complex
// number c. The return value will be floating point type corresponding to
// the type of c.

// imag �ڽ��������ظ��� c ���鲿��
// �䷵��ֵΪ��Ӧ�� c ���͵ĸ�������
func imag(c ComplexType) FloatType

// The close built-in function closes a channel, which must be either
// bidirectional or send-only. It should be executed only by the sender,
// never the receiver, and has the effect of shutting down the channel after
// the last sent value is received. After the last value has been received
// from a closed channel c, any receive from c will succeed without
// blocking, returning the zero value for the channel element. The form
//	x, ok := <-c
// will also set ok to false for a closed channel.

// close �ڽ������ر��ŵ������ŵ�����Ϊ˫��Ļ�ֻ���͵ġ�
// ��Ӧ��ֻ�ɷ�����ִ�У�����Ӧ�ɽ�����ִ�У���Ч����������͵�ֵ�����պ�ֹͣ���ŵ���
// �����һ��ֵ���ѹرյ��ŵ� c �б����պ��κδ� c �Ľ��ղ��������������ɹ���
// ���᷵�ظ��ŵ�Ԫ�����͵���ֵ�������ѹرյ��ŵ�����ʽ
//	x, ok := <-c
// ���Ὣ ok ��Ϊ false��
func close(c chan<- Type)

// The panic built-in function stops normal execution of the current
// goroutine. When a function F calls panic, normal execution of F stops
// immediately. Any functions whose execution was deferred by F are run in
// the usual way, and then F returns to its caller. To the caller G, the
// invocation of F then behaves like a call to panic, terminating G's
// execution and running any deferred functions. This continues until all
// functions in the executing goroutine have stopped, in reverse order. At
// that point, the program is terminated and the error condition is reported,
// including the value of the argument to panic. This termination sequence
// is called panicking and can be controlled by the built-in function
// recover.

// panic �ڽ�����ֹͣ��ǰGo�̵�����ִ�С�
// ������ F ���� panic ʱ��F ������ִ�оͻ�����ֹͣ���κ��� F �Ƴٵĺ���ִ�ж���
// ����һ��ķ�ʽ���У����� F ���ظ�������ߡ������������ G��F ��������Ϊ��ͬ
// �� panic �ĵ��ã�����ֹ G ��ִ�в������κα��Ƴٵĺ���������������Go��
// �����к��������෴��˳��ִֹͣ��֮�󡣴�ʱ���ó���ᱻ��ֹ������������ᱻ���棬
// ���������ÿֻŵ�ʵ��ֵ������ֹ���г�Ϊ�ֻŹ��̣�����ͨ���ڽ����� recover ���ơ�
func panic(v interface{})

// The recover built-in function allows a program to manage behavior of a
// panicking goroutine. Executing a call to recover inside a deferred
// function (but not any function called by it) stops the panicking sequence
// by restoring normal execution and retrieves the error value passed to the
// call of panic. If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the goroutine is not
// panicking, or if the argument supplied to panic was nil, recover returns
// nil. Thus the return value from recover reports whether the goroutine is
// panicking.

// recover �ڽ�����������������ֻŹ����е�Go�̡�
// �����Ƴٺ������������κα������õĺ������У�ִ�� recover ���û�ͨ���ָ�������ִ��
// ��ȡ�ش��� panic ���õĴ���ֵ��ֹͣ�ÿֻŹ������С��� recover �����Ƴٺ���֮�ⱻ���ã�
// ��������ֹͣ�ֻŹ������С��ڴ�����£��򵱸�Go�̲��ڿֻŹ�����ʱ�����ṩ�� panic
// ��ʵ��Ϊ nil ʱ��recover �ͻ᷵�� nil����� recover �ķ���ֵ�ͱ����˸�Go���Ƿ�
// �ڿֻŹ����С�
func recover() interface{}

// The print built-in function formats its arguments in an implementation-
// specific way and writes the result to standard error.
// Print is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func print(args ...Type)

// The println built-in function formats its arguments in an implementation-
// specific way and writes the result to standard error.
// Spaces are always added between arguments and a newline is appended.
// Println is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func println(args ...Type)

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.

// error �ڽ��ӿ������Ǳ�ʾ���������Լ���ӿڣ�nil ֵ����ʾû�д���
type error interface