window.onload = () => {
    // Get all images with the data-src attribute
    const images = document.querySelectorAll("img[data-src]");
    images.forEach((img) => {
        // Set the src attribute to the data-src value to trigger loading
        img.src = img.getAttribute("data-src");
        img.removeAttribute("data-src"); // Clean up
    });
    // create configuration object
    var config = {
      container: document.getElementById('heatmapContainer'),
    };
    // create heatmap with configuration
    var heatmapInstance = h337.create(config);
    // a single datapoint
    // now generate some random data
    var points = [];
    var max = 0;
    var len = 200
    var body = document.body,
    html = document.documentElement;;
    var height = Math.max( body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight );
    var width = Math.max( body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight );

    //cool while loop
    while (len--) {
      var val = Math.floor(Math.random()*100);
      max = Math.max(max, val);
      var point = {
        x: Math.floor(Math.random()*width),
        y: Math.floor(Math.random()*height),
        value: val
      };
      points.push(point);
    }
    // heatmap data format
    var data = {
      max: max,
      data: points
    };
    // if you have a set of datapoints always use setData instead of addData
    // for data initialization
    heatmapInstance.setData(data);
}
