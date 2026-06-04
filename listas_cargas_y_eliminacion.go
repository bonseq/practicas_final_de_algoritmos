carga apilada
prim:=nil

mientras p.date<>"marca fin" hacer
	nuevo (p)
	p.dato:=dato
	p.prox:=prim
	prim:=p
	leer dato
fm


leer(valor)
prim:= nil
aux:=prim
mientras p.dat<>"marca fin" hacer
	nuevo(p) //creo el nodo
	p.dato:=dato //almaceno dato
	p.prox:=nil//(apunta a nil) els iguiente
	si prim=nil entonces //si prim igual a nill significa que la lista esta vacia
		prim:=p //hay un unico elemento 
		aux:=p //esto por si quiero volver al principio de la lisat

	sino  //si ya se cargaron elementos 
		a.prox:=p
	fs
	a:=p
	leer(dato)

carga ordenada
prim:=nil
aux:=nil

mientras p<>nil p.dato<valor
	a:=q
	q:=q.prox
fm

si a=nil
	p.prox:=prim
	prim:=p
sino
	a.prox:=p
	p.prox:=q
fs


inserciones

ACCION INSERCION_ORDENADA (PRIM: PUNTERO A NODO; valor: ENTERO)
AMBIENTE
    NODO = REGISTRO
        DATO: Entero;
        prox: Puntero a NODO;
    FIN REGISTRO;
    P, Q, ANT: Puntero a NODO;

PROCESO
    NUEVO(Q); // Pedimos memoria para el nuevo p
    *Q.DATO := valor;
    
    SI (PRIM = NIL) ENTONCES
        // CASO 1: Lista Vacía [cite: 581-582]
        PRIM := Q;
        *Q.prox := NIL;
    SINO
        P := PRIM;
        ANT := NIL;
        // Buscamos el lugar correcto [cite: 584-589]
        MIENTRAS (P <> NIL) ^ (*P.DATO < *Q.DATO) HACER
            ANT := P;
            P := *P.prox;
        FIN MIENTRAS;

        SI (P = PRIM) ENTONCES
            // CASO 2: Al principio [cite: 590-591]
            *Q.prox := PRIM;
            PRIM := Q;
        SINO
            // CASO 3 y 4: En el medio o al final [cite: 592-593]
            *Q.prox := P; 
            *ANT.prox := Q;
        FIN SI;
    FIN SI;
FIN ACCION;

ACCION INSERTAR_DOBLE (REF PRIM, ULT: PUNTERO A NODO; valor: Entero)
PROCESO
    NUEVO(Q);
    *Q.DATO := valor;
    *Q.ant := NIL;
    *Q.prox := NIL;

    SI (PRIM = NIL) ENTONCES
        PRIM := Q;
        ULT := Q;
    SINO
        P := PRIM;
        MIENTRAS (P <> NIL) ^ (*P.DATO < *Q.DATO) HACER
            P := *P.prox;
        FIN MIENTRAS;

        SI (P = PRIM) ENTONCES
            *Q.prox := PRIM;
            *PRIM.ant := Q;
            PRIM := Q;
        SINO SI (P = NIL) ENTONCES
            *Q.ant := ULT;
            *ULT.prox := Q;
            ULT := Q;
        SINO
            *Q.prox := P;
            *Q.ant := *P.ant;
            *(*P.ant).prox := Q;
            *P.ant := Q;
        FIN SI;
    FIN SI; 
FIN ACCION;




PROCEDIMIENTO CARGA_IAF() ES
    AMBIENTE 

    PROCESO
        NUEVO(Q); // Pedimos memoria para el nuevo p
        *Q.CLIENTE_ID := R.CLIENTE_ID;
        *Q.TOTAL:=RESGUARDO_TOTAL_CLIENTE
        *Q.CAMT_OPERADP:=OPERACIONES
        *Q.IAF:=TOTAL_CLIENTE*OPERACIONES
        
        SI (PRIM = NIL) ENTONCES
            // LISTA VACIA
            PRIM := Q;
            *Q.prox := NIL;
        SINO
            P := PRIM;
            ANT := NIL;
            // BUSCARLUGAR
            MIENTRAS (P <> NIL) HACER
                SI (*P.IAF < *Q.IAF) ENTONCES
                    ANT := P;
                    P := *P.prox;
                SINO 
                    SI o (*p.CAMT_OPERADP<*Q.CAMT_OPERADP)
            FIN MIENTRAS;

            SI (P = PRIM) ENTONCES
                // si es al principio
                *Q.prox := PRIM;
                PRIM := Q;
            SINO
                // en el medio o al final
                *Q.prox := P; 
                *ANT.prox := Q;
            FIN SI;
        FIN SI;


    FP
FP


CARGA ord  ORDENADA 
PRIM:=NIL
A:=NIL
LEER(VALOR)
NUEVO (P)
*P.DATO;=VALOR
PUNTERODELISTA:=PRIM

MIENTRAS PUNTERODELISTA<>NIL Y *PUNTERODELISTA.DATO<VALOR //ESTO DE ACA PUEDE CAMBIAR DEPENDIENDO DE SI LA CARGA ES DE MENOR A MAYOR O DE MAYOR A MENOR
    A:=PUNTERODELISTA
    PUNTERODELISTA:=PUNTERODELISTA.prox
FIN MIENTRAS

SI A=NIL ENTONCES
    PRIM:=p
    P.prox:=NIL
SINO
    A.PROX:=p
    P.PROX:=PUNTERODELISTA
FS


ELIMINACION




ALGORITMO Carga_Doble_Circular_Ordenada
AMBIENTE
    NODO = REGISTRO
        DATO: Entero;
        prox, ant: PUNTERO A NODO;
    FIN REGISTRO;
    
    PRIM, P, Q: PUNTERO A NODO;
    valor: Entero;

PROCESO
    PRIM := NIL;
    
    LEER(valor);
    MIENTRAS (valor <> 0) HACER
        NUEVO(Q);
        *Q.DATO := valor;
        
        SI (PRIM = NIL) ENTONCES
            // CASO A: Lista vacía
            PRIM := Q;
            *Q.prox := Q;
            *Q.ant := Q;
        SINO
            // CASO B: Inserción al principio (el nuevo es el más chico)
            SI (*Q.DATO < *PRIM.DATO) ENTONCES
                *Q.prox := PRIM;
                *Q.ant := *PRIM.ant;
                *(*PRIM.ant).prox := Q;
                *PRIM.ant := Q;
                PRIM := Q; // <- La única diferencia: avisamos que hay un nuevo PRIMERO
            SINO
                // CASO C: Búsqueda del lugar (Medio o Final)
                P := *PRIM.prox; // Empezamos a mirar desde el segundo nodo
                
                // Mientras no demos la vuelta completa Y el dato sea mayor... avanzamos
                MIENTRAS (P <> PRIM) Y (*P.DATO < *Q.DATO) HACER
                    P := *P.prox;
                FIN MIENTRAS;
                
                // Insertamos ANTES del nodo P donde frenó el bucle
                *Q.prox := P;
                *Q.ant := *P.ant;
                *(*P.ant).prox := Q;
                *P.ant := Q;
            FIN SI;
        FIN SI;
        
        LEER(valor);
    FIN MIENTRAS;
FIN PROCESO


ALGORITMO Carga_Doble_Lineal_Ordenada
AMBIENTE
    NODO = REGISTRO
        DATO: Entero;
        prox, ant: PUNTERO A NODO;
    FIN REGISTRO;
    
    PRIM, ULT, P, Q: PUNTERO A NODO;
    valor: Entero;

PROCESO
    PRIM := NIL;
    ULT := NIL;
    
    LEER(valor);
    MIENTRAS (valor <> 0) HACER
        NUEVO(Q);
        *Q.DATO := valor;
        
        SI (PRIM = NIL) ENTONCES
            // CASO A: Lista vacía
            *Q.prox := NIL;
            *Q.ant := NIL;
            PRIM := Q;
            ULT := Q;
        SINO
            // Búsqueda del lugar
            P := PRIM;
            MIENTRAS (P <> NIL) Y (*P.DATO < *Q.DATO) HACER
                P := *P.prox;
            FIN MIENTRAS;
            
            SI (P = PRIM) ENTONCES
                // CASO B: Al principio (El nuevo es el más chico)
                *Q.prox := PRIM;
                *Q.ant := NIL;
                *PRIM.ant := Q;
                PRIM := Q;
            SINO SI (P = NIL) ENTONCES
                // CASO C: Al final (El bucle se cayó de la lista porque el nuevo es el más grande)
                *Q.prox := NIL;
                *Q.ant := ULT;
                *ULT.prox := Q;
                ULT := Q;
            SINO
                // CASO D: En el medio (P se frenó en un nodo mayor al nuevo)
                *Q.prox := P;
                *Q.ant := *P.ant;
                *(*P.ant).prox := Q; // El de atrás ahora apunta a Q
                *P.ant := Q;         // P ahora apunta hacia atrás a Q
            FIN SI;
        FIN SI;
        
        LEER(valor);
    FIN MIENTRAS;
FIN PROCESO

ALGORITMO Carga_Doble_Lineal_Al_Final
AMBIENTE
    NODO = REGISTRO
        DATO: Entero;
        prox, ant: PUNTERO A NODO;
    FIN REGISTRO;
    
    PRIM, ULT, Q: PUNTERO A NODO;
    valor: Entero;

PROCESO
    // Inicializamos AMBOS punteros externos
    PRIM := NIL;
    ULT := NIL;
    
    LEER(valor);
    MIENTRAS (valor <> 0) HACER
        NUEVO(Q);
        *Q.DATO := valor;
        *Q.prox := NIL; // Como va al final, siempre apuntará a NIL hacia adelante
        
        SI (PRIM = NIL) ENTONCES
            // CASO 1: Lista Vacía
            *Q.ant := NIL; // Es el primero, no hay nadie atrás
            PRIM := Q;
            ULT := Q;
        SINO
            // CASO 2: Inserción al final usando ULT
            *Q.ant := ULT;       // El nuevo mira hacia atrás al que era el último
            *ULT.prox := Q;      // El viejo último mira hacia adelante al nuevo
            ULT := Q;            // Actualizamos el puntero externo al nuevo final
        FIN SI;
        
        LEER(valor);
    FIN MIENTRAS;
FIN PROCESO