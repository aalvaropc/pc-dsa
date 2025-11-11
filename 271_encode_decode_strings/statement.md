# 🔐 271. Encode and Decode Strings / Codificar y Decodificar Cadenas

## 🇬🇧 English

**Problem:**  
Design an algorithm to **encode** a list of strings into a single string.  
The encoded string is then sent over a network and must be **decoded** back into the original list of strings.

You need to implement two functions:
- `encode(strs: List[str]) -> str`
- `decode(s: str) -> List[str]`

---

### Example
Input: ["lint", "code", "love", "you"]
Output: ["lint", "code", "love", "you"]
Explanation:
One possible encoding method is: "4#lint4#code4#love3#you"


In this encoding scheme:
- Each string is prefixed with its length and a separator (e.g. `#`),  
  so during decoding you can extract the exact substring.

---

### Constraints
- The number of strings is in the range [0, 10⁴].
- The length of each string is in the range [0, 10⁴].
- Each string consists of any possible valid ASCII characters.
- The overall data size fits within a 32-bit integer.

---

## 🇪🇸 Español

**Problema:**  
Diseña un algoritmo para **codificar** una lista de cadenas en una sola cadena.  
La cadena codificada se envía por la red y luego debe **decodificarse** de nuevo a la lista original.

Debes implementar dos funciones:
- `encode(strs: List[str]) -> str`
- `decode(s: str) -> List[str]`

---

### Ejemplo
Entrada: ["lint", "code", "love", "you"]
Salida: ["lint", "code", "love", "you"]
Explicación:
Un método posible de codificación es: "4#lint4#code4#love3#you"

En este método:
- Cada cadena se precede por su longitud y un separador (por ejemplo, `#`),  
  lo que permite al decodificar saber cuántos caracteres leer.

---

### Restricciones
- El número de cadenas está en el rango [0, 10⁴].  
- La longitud de cada cadena está en el rango [0, 10⁴].  
- Cada cadena puede contener cualquier carácter ASCII válido.  
- El tamaño total de los datos cabe en un entero de 32 bits.




