# 🧩 36. Valid Sudoku / Sudoku Válido

## 🇬🇧 English

**Problem:**  
Determine if a **9 × 9** Sudoku board is **valid**. Only the filled cells need to be validated according to the rules:

1. Each **row** must contain the digits **1–9** without repetition.  
2. Each **column** must contain the digits **1–9** without repetition.  
3. Each of the nine **3 × 3 sub-boxes** must contain the digits **1–9** without repetition.

**Notes:**
- A partially filled Sudoku board can be valid but not necessarily solvable.
- Only filled cells must follow the rules above.

---

### Example 1
Input:
-
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
-
Output: true

### Example 2
Input:
-
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
-
Output: false
Explanation:
Two '8's appear in the top-left 3×3 sub-box, so the board is invalid.

---

### Constraints
- `board.length == 9`  
- `board[i].length == 9`  
- `board[i][j]` is a digit `'1'`–`'9'` or `'.'`.

---

## 🇪🇸 Español

**Problema:**  
Determina si un tablero de Sudoku de **9 × 9** es **válido**. Solo se validan las celdas llenas según estas reglas:

1. Cada **fila** debe contener los dígitos **1–9** sin repetición.  
2. Cada **columna** debe contener los dígitos **1–9** sin repetición.  
3. Cada uno de los nueve **subcuadros de 3 × 3** debe contener los dígitos **1–9** sin repetición.

**Notas:**
- Un tablero parcialmente lleno puede ser válido, pero no necesariamente resoluble.
- Solo las celdas llenas deben cumplir las reglas.

---

### Ejemplo 1
Entrada:
-
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
-
Salida: true


### Ejemplo 2
Entrada:
-
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
-
Salida: false
Explicación:
Hay dos '8' en el subcuadro superior izquierdo de 3×3, por lo que el tablero es inválido.


---

### Restricciones
- `board.length == 9`  
- `board[i].length == 9`  
- `board[i][j]` es un dígito `'1'`–`'9'` o `'.'`.