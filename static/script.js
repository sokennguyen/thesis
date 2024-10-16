//TimeMe is a time measuring library included in html
window.onload = function() {
    TimeMe.initialize({
        currentPageName: "my-home-page", // current page
        idleTimeoutInSeconds: 5, 
    });  
    TimeMe.trackTimeOnElement('hero');
}

function getIdFromParam() {
    
};

function sendPostRequest(event) {
    xmlhttp=new XMLHttpRequest();
    xmlhttp.open("POST","/",true);
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    let timeSpentOnPage = TimeMe.getTimeOnCurrentPageInSeconds();
    let heroTime = TimeMe.getTimeOnElementInSeconds('hero');
    let returnObj = {
        "top": timeSpentOnPage,
        "hero": heroTime,
    }
    let stringReturn = JSON.stringify(returnObj)
    xmlhttp.send(stringReturn);
};

const intervalId = setInterval(sendPostRequest, 1000);

