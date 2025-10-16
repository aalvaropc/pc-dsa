from typing import List


class Problem:
    def calculate(self, nums: List[int], target: int) -> List[int]:
        nmap: dict[int, int] = {}
        for i, num in enumerate(nums):
            complement = target - num
            if complement in nmap:
                return [i, nmap[complement]]
            nmap[num] = i
        return []

if __name__ == '__main__':
    problem = Problem()

    nums = [2,7,11,15]
    target = 9

    result = problem.calculate(nums, target)
    print(result)