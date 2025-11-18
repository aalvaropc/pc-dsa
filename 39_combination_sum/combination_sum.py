from typing import List

class Solution:
    def combination_sum(self, candidates: List[int], target: int) -> List[List[int]]:
        res: List[List[int]] = []
        path: List[int] = []

        def backtrack(start: int, remaining: int) -> None:
            if remaining == 0:
                res.append(path.copy())
                return
            if remaining < 0:
                return

            for i in range(start, len(candidates)):
                path.append(candidates[i])
                backtrack(i, remaining - candidates[i])
                path.pop()

        backtrack(0, target)
        return res

        

if __name__ == "__main__":
    candidates = [2,3,6,7]
    target = 7
    s = Solution()
    result = s.combination_sum(candidates, target)
    print(result)
    