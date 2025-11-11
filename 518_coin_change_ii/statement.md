# 518. Coin Change II / Cambio de Monedas II

## 🇬🇧 English

**Problem:**  
You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the **number of combinations** that make up that amount.  
If that amount of money cannot be made up by any combination of the coins, return `0`.

You may assume that you have an **infinite number** of each kind of coin.

The answer is guaranteed to fit into a signed **32-bit integer**.

---

### Example 1
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: There are four ways to make up the amount:

- 5 = 5
- 5 = 2 + 2 + 1
- 5 = 2 + 1 + 1 + 1
- 5 = 1 + 1 + 1 + 1 + 1


### Example 2
Input: amount = 3, coins = [2]
Output: 0
Explanation: The amount 3 cannot be made up just with coins of 2.

### Example 3
Input: amount = 10, coins = [10]
Output: 1

---

### Constraints
- 1 <= coins.length <= 300  
- 1 <= coins[i] <= 5000  
- All the values of `coins` are unique.  
- 0 <= amount <= 5000  

---

## 🇪🇸 Español

**Problema:**  
Se te da un arreglo de enteros `coins` que representa monedas de diferentes denominaciones y un entero `amount` que representa una cantidad total de dinero.

Devuelve el **número de combinaciones** que forman esa cantidad.  
Si no es posible formar el monto con las monedas dadas, devuelve `0`.

Puedes suponer que tienes una **cantidad infinita** de cada tipo de moneda.

La respuesta está garantizada para caber en un entero **con signo de 32 bits**.

---

### Ejemplo 1
Entrada: amount = 5, coins = [1,2,5]
Salida: 4
Explicación: Hay cuatro formas de formar la cantidad:

- 5 = 5
- 5 = 2 + 2 + 1
- 5 = 2 + 1 + 1 + 1
- 5 = 1 + 1 + 1 + 1 + 1

### Ejemplo 2
Entrada: amount = 3, coins = [2]
Salida: 0
Explicación: La cantidad 3 no puede formarse solo con monedas de 2.

### Ejemplo 3
Entrada: amount = 10, coins = [10]
Salida: 1

---

### Restricciones
- 1 <= coins.length <= 300  
- 1 <= coins[i] <= 5000  
- Todos los valores de `coins` son únicos.  
- 0 <= amount <= 5000  
