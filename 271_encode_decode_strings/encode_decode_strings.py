from typing import List

class Solution:
    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            res += str(len(s)) + "#" + s
        
        return res

    def decode(self, s: str) -> List[str]:
        res = []
        i = 0
        n = len(s)
        while i < n:
            j = i
            while j < n and s[j] != '#':
                j += 1
            length = int(s[i:j])
            j += 1
            word = s[j:j+length]
            res.append(word)
            i = j + length
        return res


if __name__ == "__main__":
    codec = Solution()
    assert codec.decode(codec.encode([])) == []
    assert codec.decode(codec.encode([""])) == [""]
    assert codec.decode(codec.encode(["", ""])) == ["", ""]
    assert codec.decode(codec.encode(["lint","code","love","you"])) == ["lint","code","love","you"]
    assert codec.decode(codec.encode(["a#b", "#", "###"])) == ["a#b", "#", "###"]
    assert codec.decode(codec.encode(["hola", "  espacios  ", "línea\nnueva"])) == ["hola", "  espacios  ", "línea\nnueva"]
