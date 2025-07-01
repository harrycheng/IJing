export const arabicNumTrans = (num: number): string => {
  const chineseNums = ['零', '一', '二', '三', '四', '五', '六', '七', '八', '九', '十'];
  
  if (num < 11) {
    return chineseNums[num];
  } else if (num < 20) {
    return '十' + chineseNums[num % 10];
  } else if (num < 100) {
    const tens = Math.floor(num / 10);
    const units = num % 10;
    let result = chineseNums[tens] + '十';
    if (units !== 0) {
      result += chineseNums[units];
    }
    return result;
  }
  
  return num.toString();
};