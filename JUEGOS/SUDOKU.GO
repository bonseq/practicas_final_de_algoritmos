program SudokuJuego;
uses crt;

type
  TTablero = array[1..9, 1..9] of integer;
  TFijo = array[1..9, 1..9] of boolean;

var
  solucion, tablero: TTablero;
  fijo: TFijo;
  jugando: boolean;
  fila, col, val: integer;

{ Procedimiento para inicializar el tablero basado en una plantilla válida }
procedure Inicializar;
var
  i, j, pos: integer;
  { Una solución de Sudoku 9x9 predefinida }
  datos: array[1..81] of integer = (
    5,3,4, 6,7,8, 9,1,2,
    6,7,2, 1,9,5, 3,4,8,
    1,9,8, 3,4,2, 5,6,7,
    
    8,5,9, 7,6,1, 4,2,3,
    4,2,6, 8,5,3, 7,9,1,
    7,1,3, 9,2,4, 8,5,6,
    
    9,6,1, 5,3,7, 2,8,4,
    2,8,7, 4,1,9, 6,3,5,
    3,4,5, 2,8,6, 1,7,9
  );
begin
  Randomize;
  pos := 1;
  for i := 1 to 9 do
  begin
    for j := 1 to 9 do
    begin
      solucion[i, j] := datos[pos];
      pos := pos + 1;
      
      { Ocultar aleatoriamente alrededor del 45% de los números }
      if Random(100) < 45 then
      begin
        tablero[i, j] := 0; { 0 representa una casilla vacía }
        fijo[i, j] := false;
      end
      else
      begin
        tablero[i, j] := solucion[i, j];
        fijo[i, j] := true; { Las casillas fijas no se pueden modificar }
      end;
    end;
  end;
end;

{ Procedimiento para dibujar la cuadrícula del Sudoku en pantalla }
procedure DibujarTablero;
var
  i, j: integer;
begin
  clrscr;
  Writeln('=== SUDOKU EN PASCAL :3 ===');
  Writeln('    1 2 3   4 5 6   7 8 9');
  Writeln('  +-------+-------+-------+');
  for i := 1 to 9 do
  begin
    Write(i, ' | ');
    for j := 1 to 9 do
    begin
      if tablero[i, j] = 0 then
        Write('. ')
      else
        Write(tablero[i, j], ' ');
        
      if (j mod 3 = 0) then Write('| ');
    end;
    Writeln;
    
    if (i mod 3 = 0) then
      Writeln('  +-------+-------+-------+');
  end;
  Writeln;
end;

{ Función que comprueba si el tablero actual coincide con la solución }
function VerificarVictoria: boolean;
var
  i, j: integer;
begin
  VerificarVictoria := true;
  for i := 1 to 9 do
    for j := 1 to 9 do
      if tablero[i, j] <> solucion[i, j] then
        VerificarVictoria := false;
end;

{ --- BLOQUE PRINCIPAL --- }
begin
  Inicializar;
  jugando := true;
  
  while jugando do
  begin
    DibujarTablero;
    
    { Comprobar si ganó }
    if VerificarVictoria then
    begin
      Writeln('¡FELICIDADES! Has resuelto el Sudoku perfectamente. :3');
      Break;
    end;
    
    { Interfaz de usuario }
    Writeln('Instrucciones: Ingresa Fila (1-9), Columna (1-9) y el Valor (1-9).');
    Writeln('>> Para salir del juego, ingresa 0 en Fila.');
    Writeln('>> Para borrar una casilla tuya, ingresa 0 en Valor.');
    Writeln('--------------------------------------------------------------');
    
    Write('Fila: '); 
    Readln(fila);
    
    { Condición de salida }
    if fila = 0 then
    begin
      jugando := false;
      Break;
    end;
    
    if (fila >= 1) and (fila <= 9) then
    begin
      Write('Columna: '); 
      Readln(col);
      
      if (col >= 1) and (col <= 9) then
      begin
        { Validar si la casilla es de la pista inicial }
        if fijo[fila, col] then
        begin
          Writeln('Error: Esa casilla es fija y no se puede modificar.');
          Writeln('Presiona Enter para continuar...');
          Readln;
        end
        else
        begin
          Write('Valor (1-9, o 0 para borrar): '); 
          Readln(val);
          
          if (val >= 0) and (val <= 9) then
            tablero[fila, col] := val
          else
          begin
            Writeln('Valor invalido. Presiona Enter para continuar...');
            Readln;
          end;
        end;
      end
      else
      begin
        Writeln('Columna invalida. Presiona Enter para continuar...');
        Readln;
      end;
    end
    else
    begin
      Writeln('Fila invalida. Presiona Enter para continuar...');
      Readln;
    end;
  end;
  
  Writeln;
  Writeln('Gracias por jugar. Presiona Enter para salir del programa.');
  Readln;
end.