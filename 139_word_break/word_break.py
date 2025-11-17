from typing import List

class Solution:
    def word_break(self, s: str, word_dict: List[str]) -> bool:
        word_set = set(word_dict)
        dp = [False] * (len(s) + 1)
        dp[0] = True

        for i in range(1, len(s) + 1):
            for j in range(i):
                if dp[j] and s[j:i] in word_set:
                    dp[i] = True
                    break

        return dp[-1]


if __name__ == "__main__":
    solution = Solution()
    s = "leetcode"
    word_dict = ["leet", "code"]
    print(solution.word_break(s, word_dict))
