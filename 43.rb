# Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.

# Note:

#     The length of both num1 and num2 is < 110.
#     Both num1 and num2 contains only digits 0-9.
#     Both num1 and num2 does not contain any leading zero.
#     You must not use any built-in BigInteger library or convert the inputs to integer directly.


# @param {String} num1
# @param {String} num2
# @return {String}
def multiply(num1, num2)

  # return '0' if any number is zero
  # reverse num1 & num2
  # do multiplication, sum value, save result back to string

  return '0' if num1 == '0' || num2 == '0'

  str1 = num1.reverse
  str2 = num2.reverse
  output = '0' * (str1.size + str2.size)

  idx = 0                       # first number index
  last_idx = 0
  str1.each_char do |i|
    carry = 0
    loc = 0                     # second number index
    str2.each_char do |j|
      result = i.to_i * j.to_i + carry + output[idx + loc].to_i # sum value
      last_idx = idx + loc                                      # location to store value
      output[last_idx] = (result % 10).to_s
      loc += 1
      carry = result / 10
    end

    # for the last result, save carry to output
    unless carry.zero?
      last_idx += 1
      output[last_idx] = (output[idx + loc].to_i + carry).to_s
    end

    idx += 1
  end
  output[0..last_idx].reverse
end

p multiply(9, 9)