# 39. Combination Sum / Combinación de Suma

## 🇬🇧 English

**Problem:**  
Given an array of **distinct integers** `candidates` and a target integer `target`,  
return a list of all **unique combinations** of `candidates` where the chosen numbers **sum to target**.  
You may return the combinations in **any order**.

The same number may be chosen from `candidates` an **unlimited number of times**.  
Two combinations are **unique** if the frequency of at least one chosen number differs.

The test cases are generated so that the total number of unique combinations is **less than 150**.

---

### Example 1
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7.
7 is also a candidate, and 7 = 7.
These are the only two valid combinations.


### Example 2
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

### Example 3
Input: candidates = [2], target = 1
Output: []

---

### Constraints
- 1 <= candidates.length <= 30  
- 2 <= candidates[i] <= 40  
- All elements of `candidates` are **distinct**.  
- 1 <= target <= 40  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de **enteros distintos** `candidates` y un número entero `target`,  
devuelve una lista de todas las **combinaciones únicas** de `candidates` cuya suma sea igual a `target`.  
Puedes devolver las combinaciones en **cualquier orden**.

El mismo número puede elegirse de `candidates` un **número ilimitado de veces**.  
Dos combinaciones son **únicas** si difieren en la frecuencia de al menos un número.

Los casos de prueba están diseñados de forma que el número total de combinaciones únicas sea **menor que 150**.

---

### Ejemplo 1
Entrada: candidates = [2,3,6,7], target = 7
Salida: [[2,2,3],[7]]
Explicación:
2 y 3 son candidatos, y 2 + 2 + 3 = 7.
7 también es candidato, y 7 = 7.
Estas son las únicas dos combinaciones válidas.

### Ejemplo 2
Entrada: candidates = [2,3,5], target = 8
Salida: [[2,2,2,2],[2,3,3],[3,5]]

### Ejemplo 3
Entrada: candidates = [2], target = 1
Salida: []

---

### Restricciones
- 1 <= candidates.length <= 30  
- 2 <= candidates[i] <= 40  
- Todos los elementos de `candidates` son **distintos**.  
- 1 <= target <= 40  