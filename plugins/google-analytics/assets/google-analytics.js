(function () {
  var measurementId = "";
  var initialized = false;

  function pagePath() {
    return window.location.pathname + window.location.search + window.location.hash;
  }

  function pageLocation() {
    return window.location.origin + pagePath();
  }

  function loadGtag() {
    var script = document.createElement("script");
    script.async = true;
    script.src = "https://www.googletagmanager.com/gtag/js?id=" + encodeURIComponent(measurementId);
    document.head.appendChild(script);
  }

  function initializeAnalytics(config) {
    var analyticsConfig = config && config.googleAnalytics;
    measurementId = analyticsConfig && analyticsConfig.measurementId;

    if (!measurementId || initialized) {
      return;
    }

    window.dataLayer = window.dataLayer || [];
    window.gtag = window.gtag || function () {
      window.dataLayer.push(arguments);
    };

    window.gtag("js", new Date());
    window.gtag("config", measurementId, {
      send_page_view: false
    });

    loadGtag();
    initialized = true;
  }

  function trackPageView() {
    if (!measurementId || !window.gtag) {
      return;
    }

    window.gtag("event", "page_view", {
      page_title: document.title,
      page_location: pageLocation(),
      page_path: pagePath()
    });
  }

  gitbook.events.bind("start", function (event, config) {
    initializeAnalytics(config);
  });

  gitbook.events.bind("page.change", function () {
    trackPageView();
  });
}());
