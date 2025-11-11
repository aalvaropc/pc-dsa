# 📊 347. Top K Frequent Elements / K Elementos Más Frecuentes

## 🇬🇧 English

**Problem:**  
Given an integer array `nums` and an integer `k`, return the **`k` most frequent elements**.  
You may return the answer in **any order**.

---

### Example 1
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]


### Example 2
Input: nums = [1], k = 1
Output: [1]

### Example 3
Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
Output: [1,2]

---

### Constraints
- 1 <= nums.length <= 10⁵  
- -10⁴ <= nums[i] <= 10⁴  
- `k` is in the range `[1, number of unique elements]`.  
- It is guaranteed that the answer is unique.

---

**Follow-up:**  
Your algorithm’s time complexity must be better than **O(n log n)**, where `n` is the array size.

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de enteros `nums` y un entero `k`, devuelve los **`k` elementos más frecuentes**.  
Puedes devolver la respuesta en **cualquier orden**.

---

### Ejemplo 1
Entrada: nums = [1,1,1,2,2,3], k = 2
Salida: [1,2]

### Ejemplo 2
Entrada: nums = [1], k = 1
Salida: [1]

### Ejemplo 3
Entrada: nums = [1,2,1,2,1,2,3,1,3,2], k = 2
Salida: [1,2]

---

### Restricciones
- 1 <= nums.length <= 10⁵  
- -10⁴ <= nums[i] <= 10⁴  
- `k` está en el rango `[1, número de elementos únicos]`.  
- Se garantiza que la respuesta es única.

---

**Pregunta adicional:**  
La complejidad temporal de tu algoritmo debe ser mejor que **O(n log n)**, donde `n` es el tamaño del arreglo.