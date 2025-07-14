import { get64Symbol, IjDivinatory, getDivinatoryDetailByName } from './ijing';

export const getDivinatory = async (name?: string) => {
  if (name) {
    const divinatoryDetail = await getDivinatoryDetailByName(name);
    return { divinatory: name, divinatoryDetail };
  }
  // No name provided, generate a new one
  const [divinatory, divinatoryDetail] = await IjDivinatory();
  return { divinatory, divinatoryDetail };
};

export const get64Symbols = async () => {
  const symbols = await get64Symbol();
  return { symbols };
};