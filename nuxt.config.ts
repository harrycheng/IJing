export default {
  // 基本配置
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE_URL || '/api' // 确保与后端基础路径一致
    }
  },
  
  // 添加i18n配置
  experimental: {
    reactivityTransform: true,
  }
}