# 15. 3Sum / Suma de Tres Números

## 🇬🇧 English

**Problem:**  
Given an integer array `nums`, return **all the triplets** `[nums[i], nums[j], nums[k]]` such that:
- `i != j`,  
- `i != k`,  
- `j != k`,  
and  
`nums[i] + nums[j] + nums[k] == 0`.

The solution set **must not contain duplicate triplets**.

---

### Example 1
- Input: nums = [-1,0,1,2,-1,-4]
- Output: [[-1,-1,2],[-1,0,1]]

Explanation:

- nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
- nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
- nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0

The distinct triplets are [-1,0,1] and [-1,-1,2].

The order of the triplets and their elements does not matter.


### Example 2
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.


### Example 3
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

---

### Constraints
- 3 <= nums.length <= 3000  
- -10⁵ <= nums[i] <= 10⁵  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de enteros `nums`, devuelve **todos los tripletes** `[nums[i], nums[j], nums[k]]` tales que:
- `i != j`,  
- `i != k`,  
- `j != k`,  
y además  
`nums[i] + nums[j] + nums[k] == 0`.

El conjunto de soluciones **no debe contener tripletes duplicados**.

---

### Ejemplo 1
- Entrada: nums = [-1,0,1,2,-1,-4]
- Salida: [[-1,-1,2],[-1,0,1]]

Explicación:

- nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
- nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
- nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0

Los tripletes distintos son [-1,0,1] y [-1,-1,2].

El orden de los tripletes o de los elementos no importa.

### Ejemplo 2
- Entrada: nums = [0,1,1]
- Salida: []
Explicación: El único triplete posible no suma 0.


### Ejemplo 3
- Entrada: nums = [0,0,0]
- Salida: [[0,0,0]]

Explicación: El único triplete posible suma 0.

---

### Restricciones
- 3 <= nums.length <= 3000  
- -10⁵ <= nums[i] <= 10⁵  




