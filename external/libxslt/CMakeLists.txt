cmake_minimum_required(VERSION 3.15)

project(libxslt1 VERSION 1.1.34 LANGUAGES C)

list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR})

include(CheckCSourceCompiles)
include(CheckFunctionExists)
include(CheckIncludeFiles)
include(CheckLibraryExists)
include(CheckSymbolExists)

INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../icu/source/common/")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../icu/source/i18n/")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../zlib/")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_BINARY_DIR}/../zlib/")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/src/include/")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_BINARY_DIR}/../libxml/")
IF (WIN32)
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/win32")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/win32/include")
ENDIF()

IF(UNIX)
IF(APPLE)
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/mac")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/mac/include")
ELSE()
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/linux")
INCLUDE_DIRECTORIES("${CMAKE_CURRENT_SOURCE_DIR}/../libxml/linux/include")
ENDIF()
ENDIF()

option(LIBXSLT_WITH_DEBUGGER "Add the debugging support" ON)

option(LIBXSLT_WITH_MEM_DEBUG "Add the memory debugging module" OFF)
option(LIBXSLT_WITH_MODULES "Add the module support" ON)
option(LIBXSLT_WITH_PROFILER "Add the profiling support" ON)

option(LIBXSLT_WITH_THREADS "Add multithread support" OFF)
set(LIBXSLT_WITH_TRIO OFF)
option(LIBXSLT_WITH_XSLT_DEBUG "Add the debugging code" ON)

foreach(VARIABLE IN ITEMS WITH_CRYPTO WITH_DEBUGGER WITH_MEM_DEBUG WITH_MODULES WITH_PROFILER WITH_PYTHON WITH_TRIO WITH_XSLT_DEBUG)
	if(LIBXSLT_${VARIABLE})
		set(${VARIABLE} 1)
	else()
		set(${VARIABLE} 0)
	endif()
endforeach()

set(LIBEXSLT_MAJOR_VERSION 0)
set(LIBEXSLT_MINOR_VERSION 8)
set(LIBEXSLT_MICRO_VERSION 20)

set(LIBEXSLT_DOTTED_VERSION ${LIBEXSLT_MAJOR_VERSION}.${LIBEXSLT_MINOR_VERSION}.${LIBEXSLT_MICRO_VERSION})
math(EXPR LIBEXSLT_VERSION_NUMBER "${LIBEXSLT_MAJOR_VERSION} * 10000 + ${LIBEXSLT_MINOR_VERSION} * 100 + ${LIBEXSLT_MICRO_VERSION}")
set(LIBEXSLT_VERSION_EXTRA "")

set(LIBXSLT_MAJOR_VERSION ${PROJECT_VERSION_MAJOR})
set(LIBXSLT_MINOR_VERSION ${PROJECT_VERSION_MINOR})
set(LIBXSLT_MICRO_VERSION ${PROJECT_VERSION_PATCH})

set(LIBXSLT_DOTTED_VERSION "${LIBXSLT_MAJOR_VERSION}.${LIBXSLT_MINOR_VERSION}.${LIBXSLT_MICRO_VERSION}")
math(EXPR LIBXSLT_VERSION_NUMBER "${LIBXSLT_MAJOR_VERSION} * 10000 + ${LIBXSLT_MINOR_VERSION} * 100 + ${LIBXSLT_MICRO_VERSION}")
set(LIBXSLT_VERSION_EXTRA "")

set(VERSION ${LIBXSLT_DOTTED_VERSION})
set(LIBEXSLT_VERSION ${LIBEXSLT_DOTTED_VERSION})

set(MODULE_EXTENSION "${CMAKE_SHARED_LIBRARY_SUFFIX}")

if(MSVC)
	configure_file(libxslt/win32config.h config.h COPYONLY)
	configure_file(libxslt/win32config.h win32config.h COPYONLY)
else()
	if(UNIX)
		set(CMAKE_REQUIRED_LIBRARIES m)
	endif()
	check_function_exists(clock_gettime HAVE_CLOCK_GETTIME)
	check_include_files(dlfcn.h HAVE_DLFCN_H)
	check_function_exists(fabs HAVE_FABS)
	check_include_files(float.h HAVE_FLOAT_H)
	check_function_exists(floor HAVE_FLOOR)
	check_function_exists(fprintf HAVE_FPRINTF)
	check_include_files(fp_class.h HAVE_FP_CLASS_H)
	check_function_exists(ftime HAVE_FTIME)
	check_library_exists(gcrypt gcry_control "gcrypt.h" HAVE_GCRYPT)
	check_function_exists(gettimeofday HAVE_GETTIMEOFDAY)
	check_function_exists(gmtime HAVE_GMTIME)
	check_function_exists(gmtime_r HAVE_GMTIME_R)
	check_include_files(ieeefp.h HAVE_IEEEFP_H)
	check_include_files(inttypes.h HAVE_INTTYPES_H)
	check_library_exists(pthread pthread_join "" HAVE_LIBPTHREAD)
	check_include_files(locale.h HAVE_LOCALE_H)
	check_function_exists(localtime HAVE_LOCALTIME)
	check_function_exists(localtime_r HAVE_LOCALTIME_R)
	check_include_files(math.h HAVE_MATH_H)
	check_include_files(memory.h HAVE_MEMORY_H)
	check_include_files(nan.h HAVE_NAN_H)
	check_function_exists(pow HAVE_POW)
	check_function_exists(printf HAVE_PRINTF)
	check_include_files(pthread.h HAVE_PTHREAD_H)
	check_function_exists(snprintf HAVE_SNPRINTF)
	check_function_exists(sprintf HAVE_SPRINTF)
	check_function_exists(sscanf HAVE_SSCANF)
	check_function_exists(stat HAVE_STAT)
	check_include_files(stdarg.h HAVE_STDARG_H)
	check_include_files(stdint.h HAVE_STDINT_H)
	check_include_files(stdlib.h HAVE_STDLIB_H)
	check_include_files(strings.h HAVE_STRINGS_H)
	check_include_files(string.h HAVE_STRING_H)
	check_function_exists(strxfrm_l HAVE_STRXFRM_L)
	check_include_files(sys/select.h HAVE_SYS_SELECT_H)
	check_include_files(sys/stat.h HAVE_SYS_STAT_H)
	check_include_files(sys/timeb.h HAVE_SYS_TIMEB_H)
	check_include_files(sys/time.h HAVE_SYS_TIME_H)
	check_include_files(sys/types.h HAVE_SYS_TYPES_H)
	check_function_exists(time HAVE_TIME)
	check_include_files(time.h HAVE_TIME_H)
	check_include_files(unistd.h HAVE_UNISTD_H)
	check_function_exists(vfprintf HAVE_VFPRINTF)
	check_function_exists(vsnprintf HAVE_VSNPRINTF)
	check_function_exists(vsprintf HAVE_VSPRINTF)
	check_include_files(xlocale.h HAVE_XLOCALE_H)
	check_function_exists(_stat HAVE__STAT)
	set(LT_OBJDIR ".libs/")
	set(PACKAGE "libxslt")
	set(PACKAGE_BUGREPORT "xml@gnome.org")
	set(PACKAGE_NAME "libxslt")
	set(PACKAGE_STRING "libxslt ${LIBXSLT_DOTTED_VERSION}")
	set(PACKAGE_TARNAME "libxslt")
	set(PACKAGE_URL "http://www.xmlsoft.org/libxslt")
	set(PACKAGE_VERSION ${LIBXSLT_DOTTED_VERSION})
	check_include_files("float.h;stdarg.h;stdlib.h;string.h" STDC_HEADERS)
	set(_ALL_SOURCE ON)
	set(_GNU_SOURCE ON)
	set(_POSIX_PTHREAD_SEMANTICS ON)
	set(_TANDEM_SOURCE ON)
	check_c_source_compiles("
		#define __EXTENSIONS__ 1
		int main() { return 0; }
	" __EXTENSIONS__)
	check_include_files(minix/config.h HAVE_MINIX_CONFIG_H)
	if(HAVE_MINIX_CONFIG_H)
		set(_MINIX ON)
		set(_POSIX_1_SOURCE ON)
		set(_POSIX_SOURCE ON)
	endif()
	configure_file(config.h.cmake.in ${CMAKE_CURRENT_BINARY_DIR}/config.h)
endif()

configure_file(libxslt/xsltconfig.h.in ${CMAKE_CURRENT_BINARY_DIR}/libxslt/xsltconfig.h)

set(
	LIBXSLT_HDRS
	libxslt/attributes.h
	libxslt/documents.h
	libxslt/extensions.h
	libxslt/extra.h
	libxslt/functions.h
	libxslt/imports.h
	libxslt/keys.h
	libxslt/namespaces.h
	libxslt/numbersInternals.h
	libxslt/pattern.h
	libxslt/preproc.h
	libxslt/security.h
	libxslt/templates.h
	libxslt/transform.h
	libxslt/variables.h
	libxslt/xslt.h
	${CMAKE_CURRENT_BINARY_DIR}/libxslt/xsltconfig.h
	libxslt/xsltexports.h
	libxslt/xsltInternals.h
	libxslt/xsltlocale.h
	libxslt/xsltutils.h
)

set(
	LIBXSLT_SRCS
	libxslt/attributes.c
	libxslt/attrvt.c
	libxslt/documents.c
	libxslt/extensions.c
	libxslt/extra.c
	libxslt/functions.c
	libxslt/imports.c
	libxslt/keys.c
	libxslt/namespaces.c
	libxslt/numbers.c
	libxslt/pattern.c
	libxslt/preproc.c
	libxslt/security.c
	libxslt/templates.c
	libxslt/transform.c
	libxslt/variables.c
	libxslt/xslt.c
	libxslt/xsltlocale.c
	libxslt/xsltutils.c
)

add_library(xslt STATIC ${LIBXSLT_HDRS} ${LIBXSLT_SRCS})


target_include_directories(
	xslt
	PUBLIC
	$<BUILD_INTERFACE:${CMAKE_CURRENT_BINARY_DIR}>
	$<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}>
	$<INSTALL_INTERFACE:$<INSTALL_PREFIX>/${CMAKE_INSTALL_INCLUDEDIR}>
)

if(UNIX)
	target_link_libraries(xslt PRIVATE m)
	set(M_LIBS "-lm")
endif()

