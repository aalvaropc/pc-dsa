from typing import List
from bisect import bisect_left

class Solution:
    def longest_increasing_subsequence(self, nums: List[int]) -> int:
        sub = []
        for x in nums:
            i = bisect_left(sub, x)
            if i == len(sub):
                sub.append(x)
            else:
                sub[i] = x
        return len(sub)


if __name__ == "__main__":
    nums = [10, 9, 2, 5, 3, 7, 101, 18]
    solution = Solution()
    print(solution.longest_increasing_subsequence(nums))
