#Best Time to Buy and Sell Stock II
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. 
However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.

```
Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
```

##ALGORITHMS
1. Iterate Through the Array:
Traverse the array and whenever you see a price that is higher than the previous day's price, add the difference to the total profit. This represents buying on the previous day and selling on the current day.

2. Calculate Profit:
Sum up all the positive differences between consecutive days to get the total profit.


