# 26. Remove Duplicates from Sorted Array / Eliminar Duplicados de un Arreglo Ordenado

## 🇬🇧 English

**Problem:**  
Given an integer array `nums` sorted in **non-decreasing order**, remove the duplicates **in-place** such that each unique element appears only once.  
The relative order of the elements should be kept the same.

Consider the number of unique elements in `nums` to be `k`.  
After removing duplicates, return the number of unique elements `k`.

The first `k` elements of `nums` should contain the unique numbers in sorted order.  
The remaining elements beyond index `k - 1` can be ignored.

---

### Example 1
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation:
Your function should return k = 2,
with the first two elements of nums being 1 and 2 respectively.
It doesn’t matter what you leave beyond the returned k (hence they are underscores).


### Example 2
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,,,,,_]
Explanation:
Your function should return k = 5,
with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It doesn’t matter what you leave beyond the returned k.

---

### Constraints
- 1 <= nums.length <= 3 × 10⁴  
- -100 <= nums[i] <= 100  
- `nums` is sorted in non-decreasing order.  

---

**Custom Judge (How your solution is tested):**
```java
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // Expected result with correct length

int k = removeDuplicates(nums); // Call your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
// If all assertions pass, your solution is accepted.
```

---

## 🇪🇸 Español
## 🇪🇸 Español

**Problema:**  
Dado un arreglo de enteros `nums` ordenado en **orden no decreciente**, elimina los duplicados **en el mismo arreglo** de forma que cada elemento único aparezca solo una vez.  
El orden relativo de los elementos debe mantenerse igual.

Considera que el número de elementos únicos en `nums` es `k`.  
Después de eliminar los duplicados, devuelve el número de elementos únicos `k`.

Los primeros `k` elementos de `nums` deben contener los números únicos en orden ascendente.  
Los elementos que queden más allá del índice `k - 1` pueden ignorarse.

---

### Ejemplo 1
Entrada: nums = [1,1,2]
Salida: 2, nums = [1,2,_]
Explicación:
Tu función debe devolver k = 2,
con los dos primeros elementos de nums siendo 1 y 2 respectivamente.
No importa lo que quede más allá del índice k (por eso se representan con guiones bajos).

### Ejemplo 2
Entrada: nums = [0,0,1,1,1,2,2,3,3,4]
Salida: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explicación:
Tu función debe devolver k = 5,
con los cinco primeros elementos de nums siendo 0, 1, 2, 3 y 4 respectivamente.
No importa lo que quede más allá del índice k.

---

### Restricciones
- 1 <= nums.length <= 3 × 10⁴  
- -100 <= nums[i] <= 100  
- `nums` está ordenado en orden no decreciente.  

---

**Verificación personalizada (cómo el juez evalúa tu código):**
```java
int[] nums = [...]; // Arreglo de entrada
int[] expectedNums = [...]; // Respuesta esperada con longitud correcta

int k = removeDuplicates(nums); // Llamada a tu función

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
// Si todas las aserciones pasan, tu solución será aceptada.
```