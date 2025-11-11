class Solution:
    def valid_anagram(self, s: str, t: str) -> bool:  
        if len(s) != len(t):
            return False

        counter = {}

        for char in s:
            counter[char] = counter.get(char, 0) + 1

        for char in t:
            if char not in counter or counter[char] == 0:
                return False
            counter[char] -= 1

        return True

if __name__ == "__main__":
    solution = Solution()
    s = "anagram"
    t = "nagaram"

    result = solution.valid_anagram(s, t)
    print(result)