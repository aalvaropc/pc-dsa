# 1365. How Many Numbers Are Smaller Than the Current Number / Cuántos Números Son Menores Que el Número Actual

## 🇬🇧 English

**Problem:**  
Given the array `nums`, for each `nums[i]` find how many numbers in the array are **smaller than it**.  
That is, for each `nums[i]`, count the number of valid indices `j` such that `j != i` and `nums[j] < nums[i]`.

Return the answer as an array.

---

### Example 1
Input: nums = [8,1,2,2,3]
Output: [4,0,1,1,3]
Explanation:

For nums[0] = 8 → there are four smaller numbers (1, 2, 2, 3).

For nums[1] = 1 → there are no smaller numbers.

For nums[2] = 2 → one smaller number (1).

For nums[3] = 2 → one smaller number (1).

For nums[4] = 3 → three smaller numbers (1, 2, 2).


### Example 2
Input: nums = [6,5,4,8]
Output: [2,1,0,3]

### Example 3
Input: nums = [7,7,7,7]
Output: [0,0,0,0]

---

### Constraints
- 2 <= nums.length <= 500  
- 0 <= nums[i] <= 100  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo `nums`, para cada `nums[i]` encuentra cuántos números en el arreglo son **menores que él**.  
Es decir, para cada `nums[i]` cuenta la cantidad de índices `j` válidos tales que `j != i` y `nums[j] < nums[i]`.

Devuelve la respuesta como un arreglo.

---

### Ejemplo 1
Entrada: nums = [8,1,2,2,3]
Salida: [4,0,1,1,3]
Explicación:

Para nums[0] = 8 → hay cuatro números menores (1, 2, 2, 3).

Para nums[1] = 1 → no hay números menores.

Para nums[2] = 2 → hay un número menor (1).

Para nums[3] = 2 → hay un número menor (1).

Para nums[4] = 3 → hay tres números menores (1, 2, 2).

### Ejemplo 2
Entrada: nums = [6,5,4,8]
Salida: [2,1,0,3]

### Ejemplo 3
Entrada: nums = [7,7,7,7]
Salida: [0,0,0,0]

---

### Restricciones
- 2 <= nums.length <= 500  
- 0 <= nums[i] <= 100  