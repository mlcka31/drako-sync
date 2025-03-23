export const convertWei = (wei: string | number): string => {
  const weiValue = typeof wei === "string" ? BigInt(wei) : BigInt(wei);
  const gweiValue = weiValue / BigInt(1e9);
  const ethValue = weiValue / BigInt(1e18);

  if (weiValue < BigInt(1e9)) {
    return `${weiValue} wei`;
  } else if (weiValue < BigInt(1e18)) {
    return `${gweiValue} gwei`;
  } else {
    return `${ethValue} eth`;
  }
};
