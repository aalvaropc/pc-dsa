class Solution:
    def numbers_of_1_bits(self, n: int) -> int:
        count = 0
        while n:
            n &= n - 1
            count += 1
        return count


if __name__ == "__main__":
    solution = Solution()
    n = 11
    print(solution.numbers_of_1_bits(n))
