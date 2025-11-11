# 🧩 56. Merge Intervals / Unir Intervalos

## 🇬🇧 English

**Problem:**  
Given an array of intervals where `intervals[i] = [startᵢ, endᵢ]`, merge all **overlapping intervals**, and return an array of the **non-overlapping intervals** that cover all the intervals in the input.

---

### Example 1
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation:
Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

### Example 2
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation:
Intervals [1,4] and [4,5] are considered overlapping.

### Example 3
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Explanation:
Intervals [1,4] and [4,7] are considered overlapping.


---

### Constraints
- 1 <= intervals.length <= 10⁴  
- intervals[i].length == 2  
- 0 <= startᵢ <= endᵢ <= 10⁴  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de intervalos donde `intervals[i] = [startᵢ, endᵢ]`, une todos los **intervalos que se superponen** y devuelve un arreglo con los **intervalos no superpuestos** que cubren todos los intervalos de la entrada.

---

### Ejemplo 1
Entrada: intervals = [[1,3],[2,6],[8,10],[15,18]]
Salida: [[1,6],[8,10],[15,18]]
Explicación:
Como los intervalos [1,3] y [2,6] se superponen, se combinan en [1,6].


### Ejemplo 2
Entrada: intervals = [[1,4],[4,5]]
Salida: [[1,5]]
Explicación:
Los intervalos [1,4] y [4,5] se consideran superpuestos.


### Ejemplo 3
Entrada: intervals = [[4,7],[1,4]]
Salida: [[1,7]]
Explicación:
Los intervalos [1,4] y [4,7] se consideran superpuestos.


---

### Restricciones
- 1 <= intervals.length <= 10⁴  
- intervals[i].length == 2  
- 0 <= startᵢ <= endᵢ <= 10⁴  




