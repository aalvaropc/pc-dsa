from typing import List

class Solution:
    def find_minimum_rotated_sorted_array(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1

        while left < right:
            mid = (left + right) // 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        return nums[left]


if __name__ == "__main__":
    nums = [4, 5, 6, 7, 0, 1, 2]
    solution = Solution()
    print(solution.find_minimum_rotated_sorted_array(nums))
