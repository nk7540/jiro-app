export const secToHour = (sec: number, options: {sec: boolean} = {sec: true}) => {
  return new Date(sec * 1000).toISOString().substr(11, options.sec ? 8 : 5);
};
