


MIENTRAS OPCION='Y' HACER
	W("INGRESE LA CANTIDAD DE PREGUNTAS");READ(QUESTION)
	PARA i:=1 HASTA QUESTION
		PRIM:=NIL
		A:=P
		P.DATO:=PREGUNTA 
		P.ANT:=A
		P.PROX:=NIL
		SI A=NIL ENTONCES
			P.ANT:=PRIM 
		SINO 
			P.ANT:=A 
		FS
	FP
	ESC("SE MOSTRARA ESTE CUESTIONARIO HASTA QUE DECIDA PARAR Y GENERAR OTRO, PARA PARAR ESCRIBA STOP")
	MIENTRAS OPCION<>"STOP" HACER
		//MOSTRAR CUESTIONARIO
		W("DESEA GENERAR PARAR?")
	FM
	SI i=20 ENTONCES 
		CONT:=CONT+1
	FS
	PRIM:=P
	PARA i:=1 HASTA QUESTION HACER
		Q:=P.PROX
		DISPONER(P)
	FP		



	//dolucion chat
ALGORITMO Examen_Cuestionarios
AMBIENTE
    // 1. Estructura del arreglo origen (ya viene cargado)
    REG_PREG = REGISTRO
        Texto: AN(50);
        Tipo: {'T', 'M'};
    FIN REGISTRO;

    // 2. Estructura del nodo para nuestro Cuestionario (Lista Doble Circular)
    NODO = REGISTRO
        Dato: REG_PREG;
        prox, ant: PUNTERO A NODO;
    FIN REGISTRO;

    Arr_Preguntas: ARREGLO [1..200] DE REG_PREG; // Dato dado
    Elegidas: ARREGLO [1..200] DE BOOLEANO;      // Truco para no repetir preguntas
    
    PRIM, Q, Actual: PUNTERO A NODO;
    
    opcion_cuest, opcion_estudiante: Caracter;
    cant_preg, i, j, pos_random, direccion, pasos, total_cuest_maximos: Entero;
    tipo_cuest: {'T', 'M'};

PROCESO
    total_cuest_maximos := 0;
    
    ESC("¿Desea generar un cuestionario? (S/N)");
    LEER(opcion_cuest);

    MIENTRAS (opcion_cuest = 'S') HACER
        
        // --- FASE 1: CONFIGURACIÓN ---
        REPETIR
            ESC("Ingrese cantidad de preguntas (10 a 20): ");
            LEER(cant_preg);
        HASTA QUE (cant_preg >= 10) Y (cant_preg <= 20);

        SI (cant_preg = 20) ENTONCES
            total_cuest_maximos := total_cuest_maximos + 1;
        FIN SI;

        ESC("¿Tipo de cuestionario? ('T' Solo Texto / 'M' Mix)");
        LEER(tipo_cuest);

        // Limpiamos el arreglo de elegidas para el nuevo cuestionario
        PARA i := 1 HASTA 200 HACER
            Elegidas[i] := Falso;
        FIN PARA;
        
        PRIM := NIL; // Vaciamos la lista
        i := 1;

        // --- FASE 2: GENERACIÓN DE LA LISTA DOBLE CIRCULAR ---
        MIENTRAS (i <= cant_preg) HACER
            pos_random := RANDOMPREGUNTA(); 

            // Control 1: Que no esté repetida
            SI (Elegidas[pos_random] = Falso) ENTONCES
                // Control 2: Que sea del tipo correcto (o Mix que acepta cualquiera)
                SI (tipo_cuest = 'M') o (Arr_Preguntas[pos_random].Tipo = 'T') ENTONCES
                    
                    Elegidas[pos_random] := Verdadero; // La marcamos para no volver a usarla
                    
                    // Creamos el Nodo
                    NUEVO(Q);
                    *Q.Dato := Arr_Preguntas[pos_random];
                    
                    // Inserción en Lista Doble Circular
                    SI (PRIM = NIL) ENTONCES
                        PRIM := Q;
                        *Q.prox := Q; // Se apunta a sí mismo
                        *Q.ant := Q;
                    SINO
                        // Inserción al final de la circular doble
                        *Q.prox := PRIM;
                        *Q.ant := *PRIM.ant;
                        *(*PRIM.ant).prox := Q;
                        *PRIM.ant := Q;
                    FIN SI;
                    
                    i := i + 1; // Solo avanzamos si logramos meter una pregunta válida
                FIN SI;
            FIN SI;
        FIN MIENTRAS;

        // --- FASE 3: SIMULACIÓN DE ESTUDIANTES ---
        ESC("Cuestionario generado. ¿Hay estudiantes para rendir? (S/N)");
        LEER(opcion_estudiante);

        MIENTRAS (opcion_estudiante = 'S') HACER
            Actual := PRIM; // Cada estudiante arranca en la pregunta 1
            ESC("--- INICIO DE EXAMEN ---");
            
            // Le mostramos tantas preguntas como hayamos configurado
            PARA j := 1 HASTA cant_preg HACER
                ESC("Pregunta: ", *Actual.Dato.Texto);
                
                // Calculamos el próximo salto
                direccion := MOVIMIENTO(); // 0 = avanzar, 1 = retroceder
                
                // Usamos MOD para asegurar que los pasos sean <= cant_preg - 1 como pide el enunciado
                pasos := (RANDOMPREGUNTA() MOD (cant_preg - 1)) + 1; 
                
                // Realizamos el salto navegando la lista
                PARA k := 1 HASTA pasos HACER
                    SI (direccion = 0) ENTONCES
                        Actual := *Actual.prox;
                    SINO
                        Actual := *Actual.ant;
                    FIN SI;
                FIN PARA;
            FIN PARA;

            ESC("¿Siguiente estudiante? (S/N)");
            LEER(opcion_estudiante);
        FIN MIENTRAS;
        
        // (Opcional) Acá deberíamos hacer un DISPONER() de toda la lista para liberar memoria

        ESC("¿Desea generar otro cuestionario nuevo? (S/N)");
        LEER(opcion_cuest);
    FIN MIENTRAS;

    ESC("Cantidad de cuestionarios generados con 20 preguntas (Máximo): ", total_cuest_maximos);
FIN PROCESO