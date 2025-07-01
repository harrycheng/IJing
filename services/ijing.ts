import { arabicNumTrans } from '@/utils/transfer';
import { readFileSync } from 'fs';
import { join } from 'path';

export const get64Symbol = () => {
  try {
    const data = readFileSync(join(process.cwd(), 'public', '64Symbols.txt'));
    return data.toString().split('\n');
  } catch (error) {
    console.error('文件读取失败！');
    return [];
  }
};

export const IjDivinatory = (): [string, string] => {
  const HALFDIVINATORY = ['地', '雷', '水', '泽', '山', '火', '风', '天'];

  const up = IjHalfDivinatory();
  const down = IjHalfDivinatory();

  const upStr = HALFDIVINATORY[up];
  const downStr = HALFDIVINATORY[down];

  const symbols64 = get64Symbol();
  for (const s of symbols64) {
    const i = s.indexOf(upStr + downStr);
    if (i !== -1) {
      const nameRune = Array.from(s);
      const divinatory = nameRune.slice(5).join('');
      
      let indexStr = nameRune.slice(0, 2).join('');
      const indexNum = parseInt(indexStr, 10);
      const indexCnStr = arabicNumTrans(indexNum);
      
      const nextNum = indexNum + 1;
      const nextCnStr = arabicNumTrans(nextNum);
      
      const ijStr = readFileSync(join(process.cwd(), 'public', 'IJing.txt')).toString();
      
      const sIndex = ijStr.indexOf('第' + indexCnStr + '卦');
      const eIndex = ijStr.indexOf('第' + nextCnStr + '卦');
      
      const detail = ijStr.slice(sIndex, eIndex === -1 ? undefined : eIndex);
      
      return [divinatory, detail];
    }
  }
  
  return ['', ''];
};

export const getDivinatoryDetailByName = (name: string): string => {
  const symbols64 = get64Symbol();
  const ijStr = readFileSync(join(process.cwd(), 'public', 'IJing.txt')).toString();

  for (const s of symbols64) {
    if (s.includes(name)) {
      const nameRune = Array.from(s);
      const indexStr = nameRune.slice(0, 2).join('');
      const indexNum = parseInt(indexStr, 10);
      const indexCnStr = arabicNumTrans(indexNum);
      
      const nextNum = indexNum + 1;
      const nextCnStr = arabicNumTrans(nextNum);
      
      const sIndex = ijStr.indexOf('第' + indexCnStr + '卦');
      const eIndex = ijStr.indexOf('第' + nextCnStr + '卦');
      
      const detail = ijStr.slice(sIndex, eIndex === -1 ? undefined : eIndex);
      return detail;
    }
  }
  return '未找到对应的卦象详情。';
};

export const IjHalfDivinatory = (): number => {
  let divinatory = 0;
  for (let i = 0; i < 3; i++) {
    const ijSymbols = Math.floor(Math.random() * 2);
    divinatory += ijSymbols << (3 - i - 1);
  }
  return divinatory;
};