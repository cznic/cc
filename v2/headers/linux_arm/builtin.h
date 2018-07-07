//  Copyright 2018 The CC Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

// +build ignore

#ifndef _BUILTIN_H_
#define _BUILTIN_H_

#include "predefined.h"

typedef struct {
	void *_[2];		// Go *[]interface{}
} *__builtin_va_list;

typedef void *__FILE_TYPE__;

__SIZE_TYPE__ __builtin_strlen(char *__s);
__UINT64_TYPE__ __builtin_bswap64(__UINT64_TYPE__ x);
char *__builtin_strcpy(char *__dest, char *__src);
double __builtin_copysign(double x, double y);
int __builtin_abs(int j);
int __builtin_ffs(int i);
int __builtin_isprint(int);
int __builtin_memcmp(void *__s1, void *__s2, __SIZE_TYPE__ __n);
int __builtin_printf(char *__format, ...);
int __builtin_sprintf(char *__s, char *__format, ...);
int __builtin_strcmp(char *__s1, char *__s2);
int __signbit(double x);
int __signbitf(float x);
void *__builtin_alloca(__SIZE_TYPE__ __size);
void *__builtin_memcpy(void *dest, const void *src, __SIZE_TYPE__ n);
void *__builtin_memset(void *s, int c, __SIZE_TYPE__ n);
#ifdef __ccgo__
void *__va_end;
void *__va_start;
#endif
void __builtin_abort(void);
void __builtin_exit(int __status);
void __builtin_trap(void);
void __register_stdfiles(void *, void *, void *, void *);

#define __attribute(x)
#define __attribute__(xyz)
#define __builtin_choose_expr(a, b, c) (a) ? (b) : (c)
#define __builtin_expect(exp, c) (exp)
#define __builtin_offsetof(st, m) ((__SIZE_TYPE__)(&((st *)0)->m))
#define __builtin_prefetch(addr, ...) (void)(addr)
#define __builtin_signbit(x) (sizeof(x) == sizeof(float) ? __signbitf(x) : sizeof (x) == sizeof(double) ? __signbit(x) : __signbitl(x))
#define __builtin_types_compatible_p(type1, type2) __builtin_types_compatible__((type1){}, (type2){})
#define __builtin_va_arg(ap, type) (type)ap
#define __builtin_va_copy(dest, src) dest = src
#ifdef __ccgo__
#define __builtin_va_end(ap) ap = __va_end
#define __builtin_va_start(ap, arg) ap = __va_start
#else
#define __builtin_va_end(ap) ap = (void*)0
#define __builtin_va_start(ap, arg) ap = (void*)-1
#endif
#define __complex__ _Complex
#define __const const
#define __extension__
#define __typeof__ typeof
#define __volatile volatile

#endif				/* _BUILTIN_H_ */
