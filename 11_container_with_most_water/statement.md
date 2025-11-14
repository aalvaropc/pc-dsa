# 11. Container With Most Water / Contenedor con Más Agua

## 🇬🇧 English

**Problem:**  
You are given an integer array `height` of length `n`.  
There are `n` vertical lines drawn such that the two endpoints of the `iᵗʰ` line are `(i, 0)` and `(i, height[i])`.

Find **two lines** that, together with the x-axis, form a container that can **store the most water**.  
Return the **maximum amount of water** a container can store.

**Note:** You may not slant the container.

---

### Example 1
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation:
The lines represented by [1,8,6,2,5,4,8,3,7]
form a container where the maximum area of water (blue section) is 49.

### Example 2
Input: height = [1,1]
Output: 1

---

### Constraints
- n == height.length  
- 2 <= n <= 10⁵  
- 0 <= height[i] <= 10⁴  

---

## 🇪🇸 Español

**Problema:**  
Se te da un arreglo de enteros `height` de longitud `n`.  
Existen `n` líneas verticales dibujadas de forma que los dos extremos de la línea `i` son `(i, 0)` y `(i, height[i])`.

Encuentra **dos líneas** que, junto con el eje x, formen un contenedor que pueda **almacenar la mayor cantidad de agua posible**.  
Devuelve la **cantidad máxima de agua** que el contenedor puede almacenar.

**Nota:** No puedes inclinar el contenedor.

---

### Ejemplo 1
Entrada: height = [1,8,6,2,5,4,8,3,7]
Salida: 49
Explicación:
Las líneas representadas por [1,8,6,2,5,4,8,3,7]
forman un contenedor donde el área máxima de agua (sección azul) es 49.

### Ejemplo 2
Entrada: height = [1,1]
Salida: 1

---

### Restricciones
- n == height.length  
- 2 <= n <= 10⁵  
- 0 <= height[i] <= 10⁴  