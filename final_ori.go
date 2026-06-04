//  encvio/envio/../envio#ncvio/envio/../envio#..#@
// cod_envio cod_suc dir peso estado

algoritmo fina estado
ambiente
DETALLE_PAQUETE=REGISTRO
	cod_envio
	cod_suc
	dir
	RIESGOSO:BOOL
	ESTADO
fin REGISTRO

riesgoso, noRiesgos: archivo de DETALLE_PAQUETE 

Nodo=REGISTRO
	paquete:DETALLE_PAQUETE
	prox:Nodo
fin REGISTRO


PROCESO
abrir_secuenca(), crear_archivos()
PRIM=NIL
mientras NFDS(sECUENCIAENVIO) HACER
	NUEVO(P)
	mientras V<> '/' HACER
		CARGAR_LISTA()
	FMientras
	SI (*P.ESTADO='T'O P.ESTADO='P') Y *P.RIESGOSO=TRUE ENTONCES
			CARGAR_ARCHIV_RIESGOS ()
	SINO 
		SI *P.ESTADO='D' Y *P.RIESGOSO=FALSE ENTONCES
			CARGAR_ARCHIV_NORIESGOS ()
		fin SI
	FIN SI
	TOTAL:=TOTAL+1
FMientras


//DOS





