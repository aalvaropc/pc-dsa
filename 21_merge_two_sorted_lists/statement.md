# 🔗 21. Merge Two Sorted Lists / Mezclar Dos Listas Enlazadas Ordenadas

## 🇬🇧 English

**Problem:**  
You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted linked list**.  
The merged list should be made by **splicing together the nodes** of the first two lists.

Return the **head** of the merged linked list.

---

### Example 1
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]


### Example 2
Input: list1 = [], list2 = []
Output: []

### Example 3
Input: list1 = [], list2 = [0]
Output: [0]

---

### Constraints
- The number of nodes in both lists is in the range [0, 50].  
- -100 <= Node.val <= 100  
- Both `list1` and `list2` are sorted in **non-decreasing order**.

---

## 🇪🇸 Español

**Problema:**  
Se te dan las cabeceras (`head`) de dos **listas enlazadas ordenadas**: `list1` y `list2`.

Fusiona ambas listas en una sola **lista enlazada ordenada**.  
La lista resultante debe formarse **conectando los nodos** de las dos listas originales.

Devuelve la **cabecera** (`head`) de la lista enlazada fusionada.

---

### Ejemplo 1
Entrada: list1 = [1,2,4], list2 = [1,3,4]
Salida: [1,1,2,3,4,4]

### Ejemplo 2
Entrada: list1 = [], list2 = []
Salida: []

### Ejemplo 3
Entrada: list1 = [], list2 = [0]
Salida: [0]

---

### Restricciones
- El número de nodos en ambas listas está en el rango [0, 50].  
- -100 <= Node.val <= 100  
- Tanto `list1` como `list2` están ordenadas en **orden no decreciente**.