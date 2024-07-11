package global

// How much money in the trading account
var AccountBalance = 10000.0

// What percentage of that balance I can tolerate losing
var LossTolerance = .02 // %2

// Maximum amount I can tolerate losing
var MaxLossPerTrade = AccountBalance * LossTolerance

// Percentage of the gap I want to take as profit
var ProfitPercent = .8 // 80%

/// API ///

var Url string = "https://something.com/path"
var ApiKey string = "ajshd9uah9h21u2g3h912g3uy891uyt723adada8f9d2a17"
var ApiKeyHeader = "x-company-something"
