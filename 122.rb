# Say you have an array for which the ith element is the price of a given stock on day i.

# Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times). However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
  return 0 if prices.size.zero?
  profit = 0
  prices.each_with_index do |p, i|
    break if prices[i + 1].nil?
    profit += (prices[i + 1] - p) if prices[i + 1] > p
  end
  profit
end

p max_profit([1, 2])