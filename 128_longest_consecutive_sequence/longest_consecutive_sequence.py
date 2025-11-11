from typing import List

class Solution:
    def longest_consecutive_sequence(self, nums: List[int]) -> int:
        if not nums:
            return 0

        num_set = set(nums)
        longest = 0

        for num in num_set:
            if num - 1 not in num_set:
                length = 1
                while num + length in num_set:
                    length += 1
                longest = max(longest, length)

        return longest


if __name__ == "__main__":
    nums = [100, 4, 200, 1, 3, 2]
    solution = Solution()
    print(solution.longest_consecutive_sequence(nums))
