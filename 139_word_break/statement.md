# 139. Word Break / Segmentación de Palabras

## 🇬🇧 English

**Problem:**  
Given a string `s` and a dictionary of strings `wordDict`,  
return `true` if `s` can be segmented into a **space-separated sequence** of one or more dictionary words.

Note that the same word in the dictionary **may be reused multiple times** in the segmentation.

---

### Example 1
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation:
Return true because "leetcode" can be segmented as "leet code".

### Example 2
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation:
Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

### Example 3
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

---

### Constraints
- 1 <= s.length <= 300  
- 1 <= wordDict.length <= 1000  
- 1 <= wordDict[i].length <= 20  
- `s` and `wordDict[i]` consist only of **lowercase English letters**.  
- All words in `wordDict` are **unique**.

---

## 🇪🇸 Español

**Problema:**  
Dada una cadena `s` y un diccionario de palabras `wordDict`,  
devuelve `true` si `s` puede dividirse en una **secuencia de palabras separadas por espacios** que estén todas en el diccionario.

Ten en cuenta que la misma palabra del diccionario **puede reutilizarse varias veces** en la segmentación.

---

### Ejemplo 1
Entrada: s = "leetcode", wordDict = ["leet","code"]
Salida: true
Explicación:
Devuelve true porque "leetcode" puede segmentarse como "leet code".

### Ejemplo 2
Entrada: s = "applepenapple", wordDict = ["apple","pen"]
Salida: true
Explicación:
Devuelve true porque "applepenapple" puede segmentarse como "apple pen apple".
Se permite reutilizar las palabras del diccionario.

### Ejemplo 3
Entrada: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Salida: false

---

### Restricciones
- 1 <= s.length <= 300  
- 1 <= wordDict.length <= 1000  
- 1 <= wordDict[i].length <= 20  
- `s` y `wordDict[i]` solo contienen **letras minúsculas inglesas**.  
- Todas las palabras en `wordDict` son **únicas**.