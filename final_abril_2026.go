//1 BATALLA NAVAL EN PSEUDOCODIGO
ALGOTIRMO
AMBIENTE
barcos=registro 
	existe:bool
	total:entero
	pos:(1, 2)
fin registro

x, y , POS:entero 
estadisticas: vector de 1..5


TABLERO:Matriz de barcos (1..10,1..10)  //J, K
JUGADOR: Matriz de TABLERO(1..5) //I
PROCESO
//0 recargamos los valores de la matriz con FALSE
para i:=1 a 5 hacer
	para j:=1 a 10 hacer
		para k:=1 a 10 hacer
			JUGADOR(i)(j,k):= FALSE
		fin para
	fin para
fin para

//1 PEDIR INGRESAR EN UN PARA EN LAS 5 MATRICES 
//LOS BARCOS DE LOS JUGADORES EN LA MATRIZ DE 10x10
//TENGO UN PROBLEMA SI JUSTO EN EL LIMITE PONE UN BARCO DE DOS o mas POSCIONES 
	para i:=1 a 5 hacer
		ESC("jugador",i, "primero se cargaran los barcos de una pocision, escriba cordenadas
		en forma de  x, luego y")
		PARA Z:=1 A 3 HACER
			R(X), R(Y)
			J.TABLERO.existe(X,Y):=TRUE
			J.TABLERO.total:= 1
			J.TABLERO.pos:=0
		FP
		w("ahora se pondran los de dos posciones, de nuevo ingrese coordenadas y poscion(horizontal(1) o verticaL(2) ")
		para z:=1 a 2 hacer 
			R(X, Y)
			R(POS)
			MIENTRAS (X<=9 Y POS:=1) O (Y<=9 Y POS=2)
				W("NO SE PUEDEN COLOCAR EN LOS LIMITES DADA LAS CONDICION DE POSICION, PORFAVOR INGRESE ALGUNA POSICION VALIDA")
				R(X, Y), R(POS)
			FIN MIENTRAS
			//SI J.TABLERO(X,Y)=FALSE Y J.TABLERO(X+1,Y)=FALSE Y J.TABLERO(x,Y+1)=FALSE ENTONCES 
			J.TABLERO.existe(X,Y):=TRUE
			SI POS=1 ENTONCES 
				J.TABLERO.existe(X+1,Y):=TRUE
			FIN SI 
			SI POS=2 ENTONCES 
				J.TABLERO.existe(X,Y+1):=TRUE
			FIN SI 
			J.TABLERO.total(I,(X,Y)):= 2
			J.TABLERO.pos(I,(X,Y)):=POS
		FP 

		W("AHORA SE PONDRA EL DE TRES POSICIONES, siguiendo la misma logica que el anterior")
		R(X, Y), R(POS)
		MIENTRAS (X<=8 Y POS:=1) O (Y<=8 Y POS=2)
			W("NO SE PUEDEN COLOCAR EN LOS LIMITES DADA LAS CONDICION DE POSICION, PORFAVOR INGRESE ALGUNA POSICION VALIDA")
			R(X, Y), R(POS)
		FIN MIENTRAS

		SI J.TABLERO.existe(X,Y)=FALSE Y J.TABLERO.existe(X+1,Y)=FALSE Y J.TABLERO.existe(X+2,Y)=FALSE ENTONCES 
			J.TABLERO.existe(X,Y):=TRUE
			SI POS=1 ENTONCES 
				J.TABLERO.existe(X+1, Y ):=TRUE ; J.TABLERO.existe(X+2, Y):=TRUE
			FIN SI
			SI POS=2 ENTONCES
				J.TABLERO.existe (X, Y+1):=TRUE: J.TABLERO.existe (X, Y+2):=TRUE
			FIN SI 
		FIN SI
		J.TABLERO.total(I,(X,Y)):=3
		J.TABLERO.pos(I,(X,Y)):=POS 
	FIN PARA
	
	//una vez ubicado comienzan los turnos, elige tablero a atacar y cordenada

	PARA Z:= 1 A 25 HACER  //TURNOS
		PARA I:=1 A 5 HACER //jugadores
			w("ingrese el jugador a atacar")
			R(TABLERO_ATACADO)
			W("INGRESE POSICION A ATACAR")
			SI JUGADORES(TABLERO_ATACADO,(X,Y))=TRUE ENTONCES
				W("FUEGO!")
				JUGADORES(TABLERO_ATACADO,(X,Y))=FALSE
				SI J.TABLERO.TOTAL(TABLERO_ATACADO, (X, Y))=1
					ESTADISTICAS(I):=ESTADISTICAS(I)+1
					W("FELICIDADES DERRUMBO UN BARCO")
				SINO 
					SI J.TABLERO.TOTAL(TABLERO_ATACADO,(X,Y))>1 ENTONCES
						J.TABLERO.TOTAL(TABLERO_ATACADO,(X,Y)):=J.TABLERO.TOTAL(TABLERO_ATACADO,(X,Y))-1
					FIN SI
				FIN SINO
			SINO
				W("AGUA!")
			FIN SINO
			W("TURNO DEL SIGUIENTE JUGADOR")
		FIN PARA
	FIN PARA
	MAYOR:=lv
	para z:=1 a 5 HACER
		SI MAYOR<ESTADISTICAS(Z) ENTONCES
			MAYOR:= ESTADISTICAS(Z)
			ganador:=Z
	FINPARA
	W("el ganador es el jugador", ganador,"con", MAYOR, "barcos")
FIN PROCESO   