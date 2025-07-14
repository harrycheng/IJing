// plugins/ijing.server.ts
export default defineNitroPlugin(async (nitroApp) => {
  nitroApp.hooks.hook('beforeHandler', async (event) => {
    try {
    const storage = useStorage('assets:server');
    const data64 = await storage.getItem('64Symbols.txt');

      // 在这里处理你的数据
      event.context.ijingData = data64; // 将数据添加到 event 上下文中
    } catch (error) {
      console.error('Error reading data:', error);
      event.context.ijingError = error;
    }
  });
});