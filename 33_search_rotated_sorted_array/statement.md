# 33. Search in Rotated Sorted Array / Buscar en un Arreglo Ordenado y Rotado

## 🇬🇧 English

**Problem:**  
There is an integer array `nums` sorted in ascending order (with distinct values).

Before being passed to your function, `nums` was **rotated** at an unknown index `k` (`1 <= k < nums.length`)  
so that the resulting array is:  
`[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`.

For example, `[0,1,2,4,5,6,7]` rotated by 3 indices becomes `[4,5,6,7,0,1,2]`.

Given the array `nums` after rotation and an integer `target`,  
return the **index of `target`** if it exists, or **-1** if it does not.

You must write an algorithm that runs in **O(log n)** time.

---

### Example 1
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

### Example 2
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

### Example 3
Input: nums = [1], target = 0
Output: -1

---

### Constraints
- 1 <= nums.length <= 5000  
- -10⁴ <= nums[i] <= 10⁴  
- All values in `nums` are **unique**.  
- `nums` is sorted in ascending order and possibly rotated.  
- -10⁴ <= target <= 10⁴  

---

## 🇪🇸 Español

**Problema:**  
Se tiene un arreglo de enteros `nums` ordenado en **forma ascendente** (con valores únicos).

Antes de ser pasado a la función, `nums` fue **rotado** en un índice desconocido `k` (`1 <= k < nums.length`),  
de modo que el arreglo resultante es:  
`[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`.

Por ejemplo, `[0,1,2,4,5,6,7]` rotado 3 posiciones se convierte en `[4,5,6,7,0,1,2]`.

Dado el arreglo `nums` después de la rotación y un entero `target`,  
devuelve el **índice de `target`** si existe, o **-1** si no está presente.

Debes escribir un algoritmo que funcione en **O(log n)** tiempo.

---

### Ejemplo 1
Entrada: nums = [4,5,6,7,0,1,2], target = 0
Salida: 4

### Ejemplo 2
Entrada: nums = [4,5,6,7,0,1,2], target = 3
Salida: -1

### Ejemplo 3
Entrada: nums = [1], target = 0
Salida: -1

---

### Restricciones
- 1 <= nums.length <= 5000  
- -10⁴ <= nums[i] <= 10⁴  
- Todos los valores en `nums` son **únicos**.  
- `nums` está ordenado de forma ascendente y posiblemente rotado.  
- -10⁴ <= target <= 10⁴  




