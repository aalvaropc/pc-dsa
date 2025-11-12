# 268. Missing Number / Número Faltante

## 🇬🇧 English

**Problem:**  
Given an array `nums` containing `n` **distinct numbers** in the range `[0, n]`,  
return the **only number missing** from the array.

---

### Example 1
Input: nums = [3,0,1]
Output: 2
Explanation:
n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
2 is the missing number since it does not appear in nums.


### Example 2
Input: nums = [0,1]
Output: 2
Explanation:
n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
2 is the missing number since it does not appear in nums.


### Example 3
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation:
n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
8 is the missing number since it does not appear in nums.


---

### Constraints
- n == nums.length  
- 1 <= n <= 10⁴  
- 0 <= nums[i] <= n  
- All numbers in `nums` are unique.  

---

**Follow-up:**  
Can you implement a solution using only **O(1)** extra space and **O(n)** runtime?

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo `nums` que contiene `n` **números distintos** en el rango `[0, n]`,  
devuelve el **único número faltante** en el arreglo.

---

### Ejemplo 1
Entrada: nums = [3,0,1]
Salida: 2
Explicación:
n = 3, ya que hay 3 números, por lo tanto, los números están en el rango [0,3].
El número 2 falta porque no aparece en el arreglo.


### Ejemplo 2
Entrada: nums = [0,1]
Salida: 2
Explicación:
n = 2, ya que hay 2 números, por lo tanto, el rango es [0,2].
El número 2 falta porque no aparece en el arreglo.


### Ejemplo 3
Entrada: nums = [9,6,4,2,3,5,7,0,1]
Salida: 8
Explicación:
n = 9, ya que hay 9 números, por lo tanto, el rango es [0,9].
El número 8 falta porque no aparece en el arreglo.


---

### Restricciones
- n == nums.length  
- 1 <= n <= 10⁴  
- 0 <= nums[i] <= n  
- Todos los números en `nums` son únicos.  

---

**Pregunta adicional:**  
¿Podrías implementar una solución que use solo **O(1)** espaci