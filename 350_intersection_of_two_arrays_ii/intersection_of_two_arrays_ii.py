from collections import Counter
from typing import List

class Solution:
    def intersection_two_arrays_ii(self, nums1: List[int], nums2: List[int]) -> List[int]:
        count1 = Counter(nums1)
        result = []

        for num in nums2:
            if count1[num] > 0:
                result.append(num)
                count1[num] -= 1

        return result


if __name__ == "__main__":
    nums1 = [4, 9, 5]
    nums2 = [9, 4, 9, 8, 4]
    solution = Solution()
    print(solution.intersection_two_arrays_ii(nums1, nums2))
