from typing import List
from collections import defaultdict

class Solution:
    def group_anagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)

        for word in strs:
            count = [0] * 26
            for ch in word:
                count[ord(ch) - ord('a')] += 1
            
            key = tuple(count)
            groups[key].append(word)

        return list(groups.values())


if __name__ == "__main__":
    solution = Solution()
    words = ["eat", "tea", "tan", "ate", "nat", "bat"]
    result = solution.group_anagrams(words)

    for v in result: print(v)
