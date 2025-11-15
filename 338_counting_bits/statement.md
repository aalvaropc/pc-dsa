# 338. Counting Bits / Contando Bits

## 🇬🇧 English

**Problem:**  
Given an integer `n`, return an array `ans` of length `n + 1` such that for each `i` (`0 <= i <= n`),  
`ans[i]` is the **number of 1's** in the **binary representation** of `i`.

---

### Example 1
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

### Example 2
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

---

### Constraints
- 0 <= n <= 10⁵  

---

### Follow-up
- It’s easy to come up with an **O(n log n)** solution.  
  Can you do it in **O(n)** time and possibly in a **single pass**?  
- Can you do it **without using built-in functions** (like `__builtin_popcount` in C++)?

---

## 🇪🇸 Español

**Problema:**  
Dado un número entero `n`, devuelve un arreglo `ans` de longitud `n + 1` tal que para cada `i` (`0 <= i <= n`),  
`ans[i]` sea el **número de bits en 1** en la **representación binaria** de `i`.

---

### Ejemplo 1
Entrada: n = 2
Salida: [0,1,1]
Explicación:
0 --> 0
1 --> 1
2 --> 10

### Ejemplo 2
Entrada: n = 5
Salida: [0,1,1,2,1,2]
Explicación:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101


---

### Restricciones
- 0 <= n <= 10⁵  

---

### Pregunta adicional
- Es fácil encontrar una solución de **O(n log n)**.  
  ¿Puedes hacerlo en **O(n)** y quizás en **una sola pasada**?  
- ¿Podrías hacerlo **sin usar funciones integradas** (como `__builtin_popcount` en C++)?