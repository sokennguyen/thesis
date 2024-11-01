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

//when DOM is loaded, load the time tracking script THEN start loading the images
//this starts the timer before the images are fully loaded
//and also allow the user to exit before the images fully loaded
window.onload = () => {
    //TimeMe is a time measuring library included in html
    //track hover time
    TimeMe.initialize({
        currentPageName: "my-home-page", // current page
        idleTimeoutInSeconds: 10, 
    });  

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

    const exitLandingLink = document.getElementById('exit')
    exitLandingLink.onclick = sendPostRequest

    // Get all images with the data-src attribute
    const images = document.querySelectorAll("img[data-src]");
    images.forEach((img) => {
        // Set the src attribute to the data-src value to trigger loading
        img.src = img.getAttribute("data-src");
        img.removeAttribute("data-src"); // Clean up
    });
}

function  sendPostRequest(event) {
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
        hoverObj["hover-"+id] = TimeMe.getTimeOnElementInSeconds(id)
    })

    let testObj = {
        "hovers": hoverObj,
        "clicks": clickCountsObj
    }
    let stringReturn = JSON.stringify(testObj)
    //switch page after finishing posting
    xmlhttp.onload = function() {
        if (xmlhttp.status >= 200 && xmlhttp.status < 300) {
            if (path.includes("second")) {
                nextPageSecondVer();
            } else {
                nextPageFirstVer();
            }
        } else {
            console.error("Request failed with status: " + xmlhttp.status);
        }
    };
    xmlhttp.send(stringReturn);
};

function nextPageFirstVer () {
    const id = window.location.search.substr(1)

    /*
    const numberedId = id.split("=")[1]
    if (numberedId % 2 !== 0){
        window.location.href = "/survey/1-1.html?"+id
    }
    else {
        window.location.href = "/survey/2-1.html?"+id

    }
    */
    window.location.href = "/survey/1-1.html?"+id
}

function nextPageSecondVer () {
    const id = window.location.search.substr(1)
    //Odd id == first-ver is showed first. Even id == second-ver is showed first
    //if second-ver is showed first, interim is followed
    /*
    const numberedId = id.split("=")[1]
    if (numberedId % 2 !== 0){
        window.location.href = "/survey/2-1.html?"+id
    }
    else {
        window.location.href = "/survey/1-1.html?"+id
    }
    */
    window.location.href = "/survey/2-1.html?"+id
}
