# Big O Notation

It's useful to have an objective way to compare the performance of an algorithm. After a first-pass solution you'll be looking for improvements. To understand when one algorithm is better than another you'll want to use Big O Notation.

Big O Notation is a measure of the time and space complexity of an algorithm. In complexity analysis we are looking at how well it scales.

## Big O Classifications

    Constant.......O(1)
    Logarithmic....O(log n)
    Linear.........O(n)
    Linearithmic...O(n log n)
    Polynomial.....O(n ^ c)
    Exponential....O(c ^ n)
    Factorial......O(n!)

## Steps

1. Identify what the input `n` is.
2. Express the max number of operations, in terms of n.
3. Eliminate all excluding the highest order terms.
4. Remove all the constant factors.

## References

- https://en.wikipedia.org/wiki/Big_O_notation
- https://www.bigocheatsheet.com
- https://www.geeksforgeeks.org/analysis-algorithms-big-o-analysis
- https://web.mit.edu/16.070/www/lecture/big_o.pdf
