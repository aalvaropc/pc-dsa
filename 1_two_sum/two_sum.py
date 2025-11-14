from typing import List, Dict


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        nmap: Dict[int, int] = {}
        for i, num in enumerate(nums):
            complement = target - num
            if complement in nmap:
                return [i, nmap[complement]]
            nmap[num] = i
        return []

if __name__ == '__main__':
    solution = Solution()

    nums = [2,7,11,15]
    target = 9

    result = solution.two_sum(nums, target)
    print(result)