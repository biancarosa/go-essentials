function analyticsIdFromEnvironment() {
  return process.env.HONKIT_GOOGLE_ANALYTICS_ID || process.env.GOOGLE_ANALYTICS_ID || "";
}

module.exports = {
  book: {
    assets: "./assets",
    js: [
      "google-analytics.js"
    ]
  },
  hooks: {
    config: function (config) {
      config.pluginsConfig = config.pluginsConfig || {};
      config.pluginsConfig.googleAnalytics = config.pluginsConfig.googleAnalytics || {};
      config.pluginsConfig.googleAnalytics.measurementId =
        config.pluginsConfig.googleAnalytics.measurementId || analyticsIdFromEnvironment();

      return config;
    }
  }
};
