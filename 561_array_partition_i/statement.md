# 561. Array Partition / Partición de Arreglo

## 🇬🇧 English

**Problem:**  
Given an integer array `nums` of `2n` integers, group these integers into `n` pairs  
`(a₁, b₁), (a₂, b₂), ..., (aₙ, bₙ)` such that the **sum of min(aᵢ, bᵢ)** for all `i` is **maximized**.  

Return the **maximum possible sum**.

---

### Example 1
Input: nums = [1,4,3,2]
Output: 4
Explanation:
All possible pairings (ignoring the order of elements) are:

(1,4), (2,3) → min(1,4) + min(2,3) = 1 + 2 = 3

(1,3), (2,4) → min(1,3) + min(2,4) = 1 + 2 = 3

(1,2), (3,4) → min(1,2) + min(3,4) = 1 + 3 = 4
Thus, the maximum possible sum is 4.

### Example 2
Input: nums = [6,2,6,5,1,2]
Output: 9
Explanation:
The optimal pairing is (2,1), (2,5), (6,6).
min(2,1) + min(2,5) + min(6,6) = 1 + 2 + 6 = 9.


---

### Constraints
- 1 <= n <= 10⁴  
- nums.length == 2 * n  
- -10⁴ <= nums[i] <= 10⁴  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de enteros `nums` de longitud `2n`, agrupa los números en `n` pares  
`(a₁, b₁), (a₂, b₂), ..., (aₙ, bₙ)` de manera que la **suma de min(aᵢ, bᵢ)** para todos los `i` sea **máxima**.  

Devuelve la **suma máxima posible**.

---

### Ejemplo 1
Entrada: nums = [1,4,3,2]
Salida: 4
Explicación:
Todas las posibles agrupaciones (ignorando el orden) son:

(1,4), (2,3) → min(1,4) + min(2,3) = 1 + 2 = 3

(1,3), (2,4) → min(1,3) + min(2,4) = 1 + 2 = 3

(1,2), (3,4) → min(1,2) + min(3,4) = 1 + 3 = 4
Por lo tanto, la suma máxima posible es 4.

### Ejemplo 2
Entrada: nums = [6,2,6,5,1,2]
Salida: 9
Explicación:
La agrupación óptima es (2,1), (2,5), (6,6).
min(2,1) + min(2,5) + min(6,6) = 1 + 2 + 6 = 9.

---

### Restricciones
- 1 <= n <= 10⁴  
- nums.length == 2 * n  
- -10⁴ <= nums[i] <= 10⁴  