/**
 * @param {number} num
 * @return {number}
 */
var findComplement = function(num) {
  var isMinus = num < 0;
  var number = isMinus ? -num : num;
  var binStr = convertToBinaryString(parseInt(number));
  var invertedStr = invertBinaryString(binStr);
  var result = convertToInt(invertedStr);

  return isMinus ? -result : result;
}

function convertToBinaryString (num) {
  var result = '';
  while(num >= 1) {
    result = (num % 2) + result;
    num = Math.floor(num / 2);
  }
  return result;
}

function invertBinaryString (str) {
  var result = '';
  for( var i = 0; i < str.length; ++i) {
    result += (str[i] == '1') ? '0' : '1';
  }
  return result;
}

function convertToInt(str) {
  var result = 0;
  for (var i = str.length - 1; i >= 0; --i) {
    result += parseInt(str[i]) * Math.pow(2, str.length - 1 - i);
  }
  return result;
}

findComplement(process.argv[2]);
