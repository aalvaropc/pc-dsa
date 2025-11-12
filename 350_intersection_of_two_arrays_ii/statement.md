# 350. Intersection of Two Arrays II / Intersección de Dos Arreglos II

## 🇬🇧 English

**Problem:**  
Given two integer arrays `nums1` and `nums2`, return an array of their **intersection**.  
Each element in the result must appear **as many times as it shows in both arrays**,  
and the result can be returned in **any order**.

---

### Example 1
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

### Example 2
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also valid.

---

### Constraints
- 1 <= nums1.length, nums2.length <= 1000  
- 0 <= nums1[i], nums2[i] <= 1000  

---

### Follow-up
- What if the given arrays are **already sorted**? How would you optimize the algorithm?  
- What if `nums1`’s size is **much smaller** than `nums2`’s size? Which algorithm is better?  
- What if elements of `nums2` are **stored on disk**, and memory is limited so you **cannot load all elements** at once?

---

## 🇪🇸 Español

**Problema:**  
Dado dos arreglos de enteros `nums1` y `nums2`, devuelve un arreglo con su **intersección**.  
Cada elemento en el resultado debe aparecer **tantas veces como aparece en ambos arreglos**,  
y el resultado puede devolverse en **cualquier orden**.

---

### Ejemplo 1
Entrada: nums1 = [1,2,2,1], nums2 = [2,2]
Salida: [2,2]

### Ejemplo 2
Entrada: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Salida: [4,9]
Explicación: [9,4] también es válido.

---

### Restricciones
- 1 <= nums1.length, nums2.length <= 1000  
- 0 <= nums1[i], nums2[i] <= 1000  

---

### Preguntas Adicionales
- ¿Qué pasa si los arreglos ya están **ordenados**? ¿Cómo optimizarías el algoritmo?  
- ¿Qué pasa si `nums1` es **mucho más pequeño** que `nums2`? ¿Qué enfoque sería mejor?  
- ¿Qué pasa si los elementos de `nums2` están **almacenados en disco