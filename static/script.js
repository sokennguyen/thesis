//TimeMe is a time measuring library included in html
window.onload = function() {
    TimeMe.initialize({
        currentPageName: "my-home-page", // current page
        idleTimeoutInSeconds: 5, 
    });  
    TimeMe.trackTimeOnElement('hero');
    TimeMe.trackTimeOnElement('feat-list');
    TimeMe.trackTimeOnElement('benefit-list');
    TimeMe.trackTimeOnElement('big-feat-1');
    TimeMe.trackTimeOnElement('big-feat-2');
    TimeMe.trackTimeOnElement('big-feat-3');
    TimeMe.trackTimeOnElement('big-feat-4');
}

function getIdFromParam() {
    
};

function sendPostRequest(event) {
    xmlhttp=new XMLHttpRequest();
    xmlhttp.open("POST","/first-ver",true);
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    let timeSpentOnPage = TimeMe.getTimeOnCurrentPageInSeconds();
    let heroTime = TimeMe.getTimeOnElementInSeconds('hero');
    let featListTime = TimeMe.getTimeOnElementInSeconds('feat-list');
    let benefitTime = TimeMe.getTimeOnElementInSeconds('benefit-list');
    let feat1Time = TimeMe.getTimeOnElementInSeconds('big-feat-1');
    let feat2Time = TimeMe.getTimeOnElementInSeconds('big-feat-2');
    let feat3Time = TimeMe.getTimeOnElementInSeconds('big-feat-3');
    let feat4Time = TimeMe.getTimeOnElementInSeconds('big-feat-4');
    let returnObj = {
        "top": timeSpentOnPage,
        "hero": heroTime,
        "featList": featListTime,
        "benList": benefitTime,
        "feat1": feat1Time,
        "feat2": feat2Time,
        "feat3": feat3Time,
        "feat4": feat4Time,
    }
    let stringReturn = JSON.stringify(returnObj)
    xmlhttp.send(stringReturn);
};

const intervalId = setInterval(sendPostRequest, 1000);

