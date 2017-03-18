// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.



#define _STRINGIFY2_(x) #x
#define _STRINGIFY_(x) _STRINGIFY2_(x)
#define GOSYM_PREFIX _STRINGIFY_(__USER_LABEL_PREFIX__)

/* This file supports C files copied from the 6g runtime library.
   This is a version of the 6g runtime.h rewritten for gccgo's version
   of the code.  */

typedef signed int   int8    __attribute__ ((mode (QI)));
typedef unsigned int uint8   __attribute__ ((mode (QI)));
typedef signed int   int16   __attribute__ ((mode (HI)));
typedef unsigned int uint16  __attribute__ ((mode (HI)));
typedef signed int   int32   __attribute__ ((mode (SI)));
typedef unsigned int uint32  __attribute__ ((mode (SI)));
typedef signed int   int64   __attribute__ ((mode (DI)));
typedef unsigned int uint64  __attribute__ ((mode (DI)));
typedef float        float32 __attribute__ ((mode (SF)));
typedef double       float64 __attribute__ ((mode (DF)));
typedef signed int   intptr __attribute__ ((mode (pointer)));
typedef unsigned int uintptr __attribute__ ((mode (pointer)));

typedef intptr		intgo; // Go's int
typedef uintptr		uintgo; // Go's uint

typedef uintptr		uintreg;

/* Defined types.  */

typedef	uint8			bool;
typedef	uint8			byte;
typedef	struct	String		String;
typedef	struct	FuncVal		FuncVal;

struct String
{
	const byte*	str;
	intgo		len;
};
struct FuncVal
{
	void	(*fn)(void);
	// variable-size, fn-specific data here
};
void	runtime_panicstring(const char*) __attribute__ ((noreturn));
