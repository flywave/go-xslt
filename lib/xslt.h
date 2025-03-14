#pragma once

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define FLYWAVE_XSLT_API __declspec(dllexport)
#else
#define FLYWAVE_XSLT_API
#endif

typedef void *xslt_stylesheet_ptr;

FLYWAVE_XSLT_API int xslt_apply_style(xslt_stylesheet_ptr style,
                                      const char *xml, char **xml_txt,
                                      size_t *xml_txt_len);

FLYWAVE_XSLT_API void xslt_free_style(xslt_stylesheet_ptr *style);

FLYWAVE_XSLT_API int xslt_make_style(const char *xsl,
                                     xslt_stylesheet_ptr *style);

FLYWAVE_XSLT_API int xslt(const char *xsl, const char *xml, char **xml_txt,
                          size_t *xml_txt_len);
