fucnion experimento(x:[1, 2 , 3]): entero es
	SI x=3 ENTONCES
		write(i)
	SINO	
		si x=2 o x=1 ENTONCES
			EXPERIMENTO:=EXPERIMENTO(RAMDOM()); i:=i+2
		FS
	FS
FIN FUNCION

i: N(5)
PROCESO 
i:=1
EXPERIMENTO(RAMDOM())