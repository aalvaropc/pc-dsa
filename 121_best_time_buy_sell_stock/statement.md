# 121. Best Time to Buy and Sell Stock / Mejor Momento para Comprar y Vender Acciones

## 🇬🇧 English

**Problem:**  
Say you have an array for which the `iᵗʰ` element represents the price of a given stock on day `i`.

If you are only permitted to complete at most **one transaction** (i.e., buy one and sell one share of the stock), design an algorithm to find the **maximum profit**.

**Note:** You cannot sell a stock before you buy one.

---

### Example 1
Input: [7,1,5,3,6,4]
Output: 5
Explanation:
Buy on day 2 (price = 1) and sell on day 5 (price = 6),
profit = 6 - 1 = 5.
Not 7 - 1 = 6, because the selling price must be after the buying price.


### Example 2
Input: [7,6,4,3,1]
Output: 0
Explanation:
No transaction is done because the price keeps decreasing.

---

### Constraints
- 1 <= prices.length <= 10⁵  
- 0 <= prices[i] <= 10⁴  

---

**Follow-up:**  
Can you solve it in **O(n)** time?

---

## 🇪🇸 Español

**Problema:**  
Supón que tienes un arreglo donde el elemento `i` representa el precio de una acción en el día `i`.

Si solo se te permite realizar **una transacción** (es decir, comprar una y vender una acción), diseña un algoritmo para encontrar la **ganancia máxima**.

**Nota:** No puedes vender una acción antes de comprarla.

---

### Ejemplo 1
Entrada: [7,1,5,3,6,4]
Salida: 5
Explicación:
Compra en el día 2 (precio = 1) y vende en el día 5 (precio = 6),
ganancia = 6 - 1 = 5.
No 7 - 1 = 6, porque el precio de venta debe ser después del de compra.

### Ejemplo 2
Entrada: [7,6,4,3,1]
Salida: 0
Explicación:
No se realiza ninguna transacción porque el precio disminuye todo el tiempo.

---

### Restricciones
- 1 <= prices.length <= 10⁵  
- 0 <= prices[i] <= 10⁴  

---

**Pregunta adicional:**  
¿Puedes resolverlo en tiempo **O(n)**?