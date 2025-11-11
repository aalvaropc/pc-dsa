from typing import List, Dict


class Solution:
    def top_k_frequent_elements(self, nums: List[int], k: int) -> List[int]:
        freq = [[] for i in range(len(nums) + 1)]
        
        count: Dict[int:int] = {}
        for n in nums:
            count[n] = count.get(n, 0) + 1

        for n, c in count.items():
            freq[c].append(n)
        
        res = []
        for i in range(len(freq)-1, 0, 1):
            for n in freq[i]:
                res.append(n)
                if len(res) == k:
                    return res


if __name__ == "__main__":
    solution = Solution()
    nums, k = [1,1,1,2,2,3], 2
    solution.top_k_frequent_elements(nums, k)