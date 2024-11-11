window.onload = async () => {
    // Get all images with the data-src attribute
    const images = document.querySelectorAll("img[data-src]");
    images.forEach((img) => {
        // Set the src attribute to the data-src value to trigger loading
        img.src = img.getAttribute("data-src");
        img.removeAttribute("data-src"); // Clean up
    });

    //get hovers data
    const id = window.location.search.substr(1)
    const response = await fetch('http://localhost:8080/avg-first-hovers?' + id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      mode: 'cors', // Ensure the request expects CORS headers
    });
    const jsonFromServer = await response.json()
    const allHovers = jsonFromServer["Hover time"] 
    console.log(allHovers) 

    //heatmap points
    let points = [];
    let max = 0;

    const sectionElements = [
      'hero',
      'feat-list',
      'benefit-list',
      'big-feat-1',
      'big-feat-2',
      'big-feat-3',
      'big-feat-4',
    ]

    const hoveredElementsNames = Object.keys(allHovers)
    //find elements' locations by hovers data fields
    hoveredElementsNames.forEach((elementName, index) => {
      const hoveredElementDOM = document.getElementById(elementName)
      //offset Top and Left
      const offsetTop = hoveredElementDOM.offsetTop;
      const offsetLeft = hoveredElementDOM.offsetLeft;
      const centerY = offsetTop + hoveredElementDOM.offsetHeight / 2;
      const centerX = offsetLeft + hoveredElementDOM.offsetWidth / 2;
      const elementHoverValue = allHovers[elementName]
      const point = {
        x: centerX,
        y: centerY,
        value: elementHoverValue * 2,
        radius: 100
      }
      if (sectionElements.includes(elementName)) {
        point["radius"] = 400
        point["value"] = point["value"] /3 
      }
      //set heatmap's max value
      max = Math.max(max, elementHoverValue);
      console.log(point)
      points.push(point)
    });
    

    // create configuration object
    var config = {
      container: document.getElementById('heatmapContainer'),
    };
    // create heatmap with configuration
    var heatmapInstance = h337.create(config);

    // heatmap data format
    var data = {
      max: max,
      data: points
    };
    heatmapInstance.setData(data);
}
