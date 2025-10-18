from typing import List

class Solution:
    def contains_duplicate(self, nums: List[int]) ->bool:
        hashset = set()
        for n in nums:
            if n in hashset:
                return True
        
            hashset.add(n)
        
        return False


if __name__ == "__main__":
    solution = Solution()
    nums = [1,2,3,1]
    result = solution.contains_duplicate(nums)
    print(result)
