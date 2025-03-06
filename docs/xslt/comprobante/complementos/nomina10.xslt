<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="2.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:fn="http://www.w3.org/2005/xpath-functions" xmlns:nomina="http://www.sat.gob.mx/nomina">
	<xsl:output method="text" version="1.0" encoding="UTF-8" indent="no"/>

	<!-- Manejador de nodos tipo nomina -->
	<xsl:template match="nomina:Nomina">
		<!--Iniciamos el tratamiento de los atributos de Nómina -->
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./@Version"/> 
		</xsl:call-template>
		
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./@RegistroPatronal"/>
		</xsl:call-template>

		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./@NumEmpleado"/>
		</xsl:call-template>
		
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./@CURP"/>
		</xsl:call-template>

		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./@TipoRegimen"/>
		</xsl:call-template>
		
		<xsl:call-template name="Opcional">
			<xsl:with-param name="valor" select="./@NumSeguridadSocial"/>
		</xsl:call-template>
		
		<xsl:call-template name="Opcional">
			<xsl:with-param name="valor" select="./@CLABE"/>
		</xsl:call-template>
		
				<xsl:call-template name="Opcional">
			<xsl:with-param name="valor" select="./@Banco"/>
		</xsl:call-template>

		<!--Iniciamos el tratamiento de los atributos de Ingresos -->
		
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./nomina:Ingresos/@TotalGravado"/>
		</xsl:call-template>
		
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./nomina:Ingresos/@TotalExento"/>
		</xsl:call-template>		
	
		<!--Iniciamos el tratamiento de los atributos de descuentos -->	
	
		<xsl:call-template name="Requerido">
			<xsl:with-param name="valor" select="./nomina:Descuentos/@Total"/>
		</xsl:call-template>		
		
	</xsl:template>
</xsl:stylesheet>