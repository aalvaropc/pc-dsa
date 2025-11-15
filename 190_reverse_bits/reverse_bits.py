class Solution:
    def reverse_bits(self, n: int) -> int:
        res = 0
        for _ in range(32):
            res = (res << 1) | (n & 1)
            n >>= 1
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.reverse_bits(43261596))
