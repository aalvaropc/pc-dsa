from typing import List

class Solution:
    def maximum_product_subarray(self, nums: List[int]) -> int:
        max_prod = min_prod = result = nums[0]

        for num in nums[1:]:
            if num < 0:
                max_prod, min_prod = min_prod, max_prod

            max_prod = max(num, num * max_prod)
            min_prod = min(num, num * min_prod)
            result = max(result, max_prod)

        return result


if __name__ == "__main__":
    nums = [2, 3, -2, 4]
    solution = Solution()
    print(solution.maximum_product_subarray(nums))
