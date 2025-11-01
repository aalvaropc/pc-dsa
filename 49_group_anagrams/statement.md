# 🔠 49. Group Anagrams / Agrupar Anagramas

## 🇬🇧 English

**Problem:**  
Given an array of strings `strs`, group the **anagrams** together.  
You can return the answer in any order.

An **anagram** is a word or phrase formed by rearranging the letters of another word or phrase, typically using all the original letters exactly once.

---

### Example 1
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explanation:

There is no string in strs that can be rearranged to form "bat".

The strings "nat" and "tan" are anagrams of each other.

The strings "ate", "eat", and "tea" are anagrams of each other.

### Example 2
Input: strs = [""]
Output: [[""]]


### Example 3
Input: strs = ["a"]
Output: [["a"]]

---

### Constraints
- 1 <= strs.length <= 10⁴  
- 0 <= strs[i].length <= 100  
- `strs[i]` consists of lowercase English letters.  

---

## 🇪🇸 Español

**Problema:**  
Dado un arreglo de cadenas `strs`, agrupa los **anagramas** entre sí.  
Puedes devolver la respuesta en cualquier orden.

Un **anagrama** es una palabra o frase formada al reorganizar las letras de otra palabra o frase, usando **todas las letras originales exactamente una vez**.

---

### Ejemplo 1
Entrada: strs = ["eat","tea","tan","ate","nat","bat"]
Salida: [["bat"],["nat","tan"],["ate","eat","tea"]]
Explicación:

No existe ninguna cadena que pueda reorganizarse para formar "bat".

Las cadenas "nat" y "tan" son anagramas entre sí.

Las cadenas "ate", "eat" y "tea" también son anagramas entre sí.

### Ejemplo 2
Entrada: strs = [""]
Salida: [[""]]

### Ejemplo 3
Entrada: strs = ["a"]
Salida: [["a"]]

---

### Restricciones
- 1 <= strs.length <= 10⁴  
- 0 <= strs[i].length <= 100  
- `strs[i]` contiene solo letras minúsculas del alfabeto inglés.  