📚 Finales de Algoritmos y Estructuras de Datos (AED)
¡Hola! Soy Taima, estudiante de Sistemas. Este repositorio es un compendio de ejercicios y resoluciones de exámenes finales de la cátedra de Algoritmos y Estructuras de Datos.

[!IMPORTANT] Nota sobre el lenguaje: Todos los algoritmos aquí presentados están escritos en pseudocódigo formal, respetando las normas y la sintaxis solicitada por la cátedra (UTN-FRRE). No es código ejecutable (como C++ o Python), sino lógica pura para aprobar el "final más difícil".

🚀 Contenido del Repositorio

LISTAS: Simples, dobles y circulares. Arboles, Recursividad 
ARCHIVOS: Indexados y Secuenciales
SECUENCIAS, ARRAYS: strings o numericos

❓ FAQ - Preguntas Frecuentes (Minitips de Supervivencia)
Aquí voy anotando esas dudas "existenciales" que surgen mientras practico para los parciales y finales.

📝 Strings y Variables
¿Cómo "pongo a cero" un String? En pseudocódigo, simplemente asignas comillas vacías: mi_string := "" o mi_string := " ".

¿0 es lo mismo que NULL? ¡No! Si un nodo contiene el valor 0, la lista no está vacía; simplemente tiene un dato numérico. Una lista está vacía cuando su puntero externo es NIL o NULL.

🔗 Punteros y Listas
Si asigno dos punteros entre sí (P := Q), ¿qué pasa? No se apuntan entre sí. Lo que sucede es que P ahora guarda la misma dirección de memoria que Q. Ambos terminan apuntando al mismo nodo.

¿Qué es un Nodo? Es un registro que contiene datos y al menos un puntero al siguiente objeto de la colección.

¿Cómo libero memoria? Siempre que elimines un elemento, debés usar la acción DISPONER(puntero) para liberar la dirección de memoria interna.

🔄 Recursividad
¿Cuándo usarla? Se debe usar cuando sea realmente necesaria o cuando no exista una solución iterativa simple.

¿Qué es el Caso Base? Es el problema trivial que se resuelve sin cálculo; la solución simple que detiene las llamadas recursivas.

🤝 ¿Cómo contribuir?
Si sos alumno y encontrás un error en la lógica o querés proponer una solución más eficiente, ¡sentite libre de abrir un Pull Request o un Issue! Todo suma para que todos aprobemos.

Hecho por Taima_pomelito 🍈
