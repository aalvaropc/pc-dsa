# 🧮 88. Merge Sorted Array / Mezclar Arreglos Ordenados

## 🇬🇧 English

**Problem:**  
You are given two integer arrays `nums1` and `nums2`, sorted in **non-decreasing order**, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

Merge `nums1` and `nums2` into a **single array sorted** in non-decreasing order.

The final sorted array **should not be returned** by the function, but instead be **stored inside `nums1`**.  
To accommodate this, `nums1` has a length of `m + n`, where:
- The first `m` elements denote the meaningful values to merge.
- The last `n` elements are `0` and should be ignored.  
`nums2` has a length of `n`.

---

### Example 1
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation:
The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6],
with the underlined elements coming from nums1.


### Example 2
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation:
The arrays we are merging are [1] and [].
The result of the merge is [1].


### Example 3
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation:
The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1.
The 0 is only a placeholder to fit the final result.


---

### Constraints
- `nums1.length == m + n`  
- `nums2.length == n`  
- 0 <= m, n <= 200  
- 1 <= m + n <= 200  
- -10⁹ <= nums1[i], nums2[j] <= 10⁹  

---

**Follow-up:**  
Can you come up with an algorithm that runs in **O(m + n)** time?

---

## 🇪🇸 Español

**Problema:**  
Se te dan dos arreglos de enteros `nums1` y `nums2`, ordenados en **orden no decreciente**, y dos enteros `m` y `n`, que representan el número de elementos en `nums1` y `nums2` respectivamente.

Debes **fusionar `nums1` y `nums2`** en un solo arreglo **ordenado** en orden no decreciente.

El arreglo final **no debe devolverse**, sino almacenarse **dentro de `nums1`**.  
Para ello:
- `nums1` tiene longitud `m + n`.  
  - Los primeros `m` elementos son los valores válidos a fusionar.  
  - Los últimos `n` son ceros (`0`) que se ignoran.  
- `nums2` tiene longitud `n`.

---

### Ejemplo 1
Entrada: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Salida: [1,2,2,3,5,6]
Explicación:
Los arreglos que se fusionan son [1,2,3] y [2,5,6].
El resultado de la fusión es [1,2,2,3,5,6],
donde los elementos subrayados provienen de nums1.

### Ejemplo 2
Entrada: nums1 = [1], m = 1, nums2 = [], n = 0
Salida: [1]
Explicación:
Los arreglos que se fusionan son [1] y [].
El resultado de la fusión es [1].

### Ejemplo 3
Entrada: nums1 = [0], m = 0, nums2 = [1], n = 1
Salida: [1]
Explicación:
Los arreglos que se fusionan son [] y [1].
El resultado de la fusión es [1].
Como m = 0, no hay elementos válidos en nums1.
El 0 solo está ahí para poder almacenar el resultado final.

---

### Restricciones
- `nums1.length == m + n`  
- `nums2.length == n`  
- 0 <= m, n <= 200  
- 1 <= m + n <= 200  
- -10⁹ <= nums1[i], nums2[j] <= 10⁹  

---

**Pregunta adicional:**  
¿Puedes idear un algoritmo que funcione en tiempo **O(m + n)**?




