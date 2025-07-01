import { get64Symbol, IjDivinatory, getDivinatoryDetailByName } from '@/services/ijing';

export const getDivinatory = async (name?: string) => {
  if (name) {
    const divinatoryDetail = getDivinatoryDetailByName(name);
    return { divinatory: name, divinatoryDetail };
  }
  // No name provided, generate a new one
  const [divinatory, divinatoryDetail] = IjDivinatory();
  return { divinatory, divinatoryDetail };
};

export const get64Symbols = async () => {
  const symbols = get64Symbol();
  return { symbols };
};