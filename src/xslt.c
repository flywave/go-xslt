#include "xslt.h"
#include <libxslt/transform.h>
#include <libxslt/xsltutils.h>
#include <stdint.h>
#include <string.h>

#include <libxslt/xsltutils.h>

int xslt_apply_style(xslt_stylesheet_ptr style, const char *xml, char **xml_txt,
                     size_t *xml_txt_len) {

  int ok;
  size_t len;
  xmlChar *xml_output;
  xmlDocPtr xml_doc, result;

  len = strlen(xml);
  if (len > INT32_MAX) {
    return -1;
  }

  xml_doc = xmlParseMemory(xml, (int)strlen(xml));
  if (xml_doc == NULL || xmlGetLastError()) {
    xmlResetLastError();
    return -1;
  }

  result = xsltApplyStylesheet((xsltStylesheetPtr)style, xml_doc, NULL);
  if (result == NULL) {
    xmlFreeDoc(xml_doc);
    return -1;
  }

  ok = xsltSaveResultToString(&xml_output, (int *)xml_txt_len, result,
                              (xsltStylesheetPtr)style);
  if (ok == 0 && *xml_txt_len > 0) {
    *xml_txt = malloc(*xml_txt_len);
    strncpy(*xml_txt, (const char *)xml_output, *xml_txt_len);
    xmlFree(xml_output);
  }

  xmlFreeDoc(xml_doc);
  xmlFreeDoc(result);

  return ok;
}

void xslt_free_style(xslt_stylesheet_ptr *style) {
  xsltFreeStylesheet((xsltStylesheetPtr)*style);
}

int xslt_make_style(const char *xsl, xslt_stylesheet_ptr *style) {

  size_t len;
  xmlDocPtr style_doc;

  len = strlen(xsl);
  if (len > INT32_MAX) {
    return -1;
  }

  style_doc = xmlParseMemory(xsl, (int)len);
  if (style_doc == NULL || xmlGetLastError()) {
    xmlResetLastError();
    return -1;
  }

  *style = xsltParseStylesheetDoc(style_doc);
  if (*style == NULL || ((xsltStylesheetPtr)(*style))->errors) {
    xmlFreeDoc(style_doc);
    return -1;
  }

  return 0;
}

int xslt(const char *xsl, const char *xml, char **xml_txt,
         size_t *xml_txt_len) {

  int ok;
  xsltStylesheetPtr style;

  ok = xslt_make_style(xsl, &style);
  if (ok < 0) {
    return -1;
  }

  ok = xslt_apply_style(style, xml, xml_txt, xml_txt_len);

  xslt_free_style(&style);

  return ok;
}
