const clickElements = [
    'nav-feat', 'nav-price', 'nav-login', 'nav-start', 
    'hero-cta', 'hero-login', 'small-feat1-pic', 
    'small-feat2-pic', 'small-feat3-pic', 'headstart', 
    'consistency','headstart', 'determination', 'big-feat1-img', 
    'big-feat2-img', 'big-feat2-cta', 'big-feat3-img', 
    'big-feat3-more','big-feat4-img', 'big-feat4-more',
    'ending-cta-btn'
];

const hoverTracks = [

    ...clickElements,
    
    //sections
    'hero',
    'feat-list',
    'benefit-list',
    'big-feat-1',
    'big-feat-2',
    'big-feat-3',
    'big-feat-4',

    //other elements
    'head-logo',
    'hero-title',
    'sub-title',
    'headstart-desc',
    'consistency-desc',
    'flexible-desc',
    'determination-desc',
    'big-feat1-desc',
    'big-feat2-desc',
    'big-feat3-desc',
    'big-feat4-desc',
    'ending-title',
    'ending-subtitle',
    'ending-cta-btn',
    'footer-logo',
    'footer-product',
    'footer-company',
    'footer-legal',
];

let clickCountsObj = {};

//TimeMe is a time measuring library included in html
window.onload = function() {
    //track hover time
    TimeMe.initialize({
        currentPageName: "my-home-page", // current page
        idleTimeoutInSeconds: 10, 
    });  

    /* 
    let clickCounts = {
        'nav-feat': 0,
        'nav-price': 0,
        'nav-login': 0,
        'nav-start': 0,
        'hero-cta': 0,
        'hero-login': 0,
        'small-feat1-img': 0,
        'small-feat2-img': 0,
        'small-feat3-img': 0,
        'headstart': 0,
        'consistency': 0,
        'determination': 0,
        'big-feat1-img': 0,
        'big-feat2-img': 0,
        'big-feat3-img': 0,
        'big-feat4-img': 0,
        'ending-cta-btn': 0,
    } 
    */




    // Initialize clickCounts object
    clickElements.forEach(id => {
        clickCountsObj[id] = 0;
    });

    const incrementOnclick = (e) => {
        const id = e.target.id 
        clickCountsObj[id] += 1;
        console.log(clickCountsObj)
    }

    clickElements.forEach(id => {
        const element = document.getElementById(id);
        if (element) {
            element.addEventListener("click", incrementOnclick);
        }
    });


    hoverTracks.forEach(id => {
        //TimeMe initialized aboved
        TimeMe.trackTimeOnElement(id);
    })
}


function sendPostRequest(event) {
    xmlhttp=new XMLHttpRequest();
    const id = window.location.search.substr(1)
    const path = location.pathname.split('/')[1]
    if (path.includes("second")){
        xmlhttp.open("POST","/second-ver?"+id,true);
    }
    else {
        xmlhttp.open("POST","/first-ver?"+id,true);
    }
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

    let timeSpentOnPage = TimeMe.getTimeOnCurrentPageInSeconds();
    let hoverObj = {"top": timeSpentOnPage}
    hoverTracks.forEach(id => {
        hoverObj[id] = TimeMe.getTimeOnElementInSeconds(id)
    })

    let testObj = {
        "hovers": hoverObj,
        "clicks": clickCountsObj
    }
    let stringReturn = JSON.stringify(testObj)
    xmlhttp.send(stringReturn);
};

const intervalId = setInterval(sendPostRequest, 1000);

