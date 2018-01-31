# @param {Integer} num
# @return {Integer}
def find_complement(num)
  minus = num.to_i < 0

  binary_string = convert_to_binary_str(num.to_i.abs)
  inverted_string = invert_binary_str(binary_string)
  value = convert_str_to_int(inverted_string)
  minus ? -value : value
end

def convert_to_binary_str(num)
  binary = ""
  while num != 0
    binary << (num % 2).to_s
    num /= 2
  end
  binary
end

def invert_binary_str(str)
  inverted = ''
  str.each_char { |c| inverted << ((c == '1') ? '0' : '1') }
  inverted
end

def convert_str_to_int(str)
  num = 0
  for i in 0..str.length
    num += str[i].to_i * 2**i
  end
  num
end

find_complement(ARGV[0])
