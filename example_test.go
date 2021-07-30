package xslt_test

import (
	"fmt"

	"github.com/flywave/go-xslt"
)

func Example() {
	var doc = []byte(
		`<?xml version="1.0" ?>
		 <persons>
		   <person username="JS1">
		     <name>John</name>
		     <family-name>Smith</family-name>
		   </person>
		   <person username="MI1">
		     <name>Morka</name>
		     <family-name>Ismincius</family-name>
		   </person>
		 </persons>`,
	)

	var style = []byte(
		`<?xml version="1.0" encoding="UTF-8"?>
		 <xsl:stylesheet
		   version="1.0"
		   xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
		   xmlns="http://www.w3.org/1999/xhtml">
		   <xsl:output method="xml" indent="yes" encoding="UTF-8"/>
		   <xsl:template match="/persons">
		     <html>
		       <head>
			 <title>Testing XML Example</title>
		       </head>
		       <body>
			 <h1>Persons</h1>
			 <ul>
			   <xsl:apply-templates select="person">
			     <xsl:sort select="family-name" />
			   </xsl:apply-templates>
			 </ul>
		       </body>
		     </html>
		   </xsl:template>
		   <xsl:template match="person">
		     <li>
		       <xsl:value-of select="family-name"/><xsl:text>, </xsl:text><xsl:value-of select="name"/>
		     </li>
		   </xsl:template>
		 </xsl:stylesheet>`,
	)

	xs, err := xslt.NewStylesheet(style)
	if err != nil {
		panic(err)
	}
	defer xs.Close()

	res, err := xs.Transform(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
