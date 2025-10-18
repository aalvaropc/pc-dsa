from typing import List

class Solution:
    def max_profit(self, prices: List[int]) -> int:
        min_prices = float('inf')
        max_profit = 0

        for price in prices:
            if price < min_prices:
                min_prices = price
            elif price - min_prices > max_profit:
                max_profit = price - min_prices
            
        return max_profit

if __name__ == '__main__':
    problem = Solution()
    prices = [7,1,5,3,6,4]
    result = problem.calculate(prices)