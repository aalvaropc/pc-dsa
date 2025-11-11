from typing import List

class Solution:
    def coin_change_ii(self, amount: int, coins: List[int]) -> int:
        dp = [0] * (amount + 1)
        dp[0] = 1

        for coin in coins:
            for x in range(coin, amount + 1):
                dp[x] += dp[x - coin]

        return dp[amount]


if __name__ == "__main__":
    solution = Solution()
    print(solution.coin_change_ii(5, [1, 2, 5]))
